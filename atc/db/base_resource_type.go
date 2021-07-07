package db

import (
	"sync"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// The table base_resource_types is small (less than 100 tuples) and steady (it
// may only be updated after Concourse changes). But it's queried extremely
// frequently, thus it makes sense to cache this table in memory, and reload it
// periodically, which should help reduce db queries.

type baseResourceTypeRaw struct {
	id                   int
	name                 string
	uniqueVersionHistory bool
}

type baseResourceTypeTableRaw struct {
	tableByName map[string]baseResourceTypeRaw

	rwLock       sync.RWMutex
	lastLoadTime time.Time
}

var (
	baseResourceTypeTable = &baseResourceTypeTableRaw{
		tableByName: map[string]baseResourceTypeRaw{},
	}
	disableBaseResourceTypeCache = false // this flag is mostly used for unit tests.
)

func (table *baseResourceTypeTableRaw) findByName(runner sq.Runner, name string) (baseResourceTypeRaw, bool, error) {
	err := table.reloadIfNeeded(runner)
	if err != nil {
		return baseResourceTypeRaw{}, false, err
	}

	table.rwLock.RLock()
	defer table.rwLock.RUnlock()

	brt, found := table.tableByName[name]
	return brt, found, nil
}

func (table *baseResourceTypeTableRaw) reloadIfNeeded(runner sq.Runner) error {
	if !disableBaseResourceTypeCache && table.lastLoadTime.Add(3 * time.Minute).After(time.Now()) {
		return nil
	}

	rows, err := psql.Select("id, name, unique_version_history").
		From("base_resource_types").
		RunWith(runner).
		Query()
	if err != nil {
		return err
	}

	defer rows.Close()

	table.rwLock.Lock()
	defer table.rwLock.Unlock()

	table.tableByName = map[string]baseResourceTypeRaw{}

	for rows.Next() {
		var row baseResourceTypeRaw
		err := rows.Scan(&row.id, &row.name, &row.uniqueVersionHistory)
		if err != nil {
			return err
		}
		table.tableByName[row.name] = row
	}
	table.lastLoadTime = time.Now()

	return nil
}

func DisableBaseResourceTypeCache() {
	disableBaseResourceTypeCache = true
}

// BaseResourceType represents a resource type provided by workers.
//
// It is created via worker registration. All creates are upserts.
//
// It is removed by gc.BaseResourceTypeCollector, once there are no references
// to it from worker_base_resource_types.
type BaseResourceType struct {
	Name string // The name of the type, e.g. 'git'.
}

// UsedBaseResourceType is created whenever a ResourceConfig is used, either
// for a build, a resource in the pipeline, or a resource type in the pipeline.
//
// So long as the UsedBaseResourceType's ID is referenced by a ResourceConfig
// that is in use, this guarantees that the BaseResourceType will not be
// removed. That is to say that its "Use" is vicarious.
type UsedBaseResourceType struct {
	ID                   int    // The ID of the BaseResourceType.
	Name                 string // The name of the type, e.g. 'git'.
	UniqueVersionHistory bool   // If set to true, will create unique version histories for each of the resources using this base resource type
}

// FindOrCreate looks for an existing BaseResourceType and creates it if it
// doesn't exist. It returns a UsedBaseResourceType.
func (brt BaseResourceType) FindOrCreate(tx Tx, unique bool) (*UsedBaseResourceType, error) {
	ubrt, found, err := brt.Find(tx)
	if err != nil {
		return nil, err
	}

	if found && ubrt.UniqueVersionHistory == unique {
		return ubrt, nil
	}

	return brt.create(tx, unique)
}

func (brt BaseResourceType) Find(runner sq.Runner) (*UsedBaseResourceType, bool, error) {
	brtRaw, found, err := baseResourceTypeTable.findByName(runner, brt.Name)
	if err != nil || !found {
		return nil, found, err
	}

	return &UsedBaseResourceType{ID: brtRaw.id, Name: brtRaw.name, UniqueVersionHistory: brtRaw.uniqueVersionHistory}, true, nil
}

func (brt BaseResourceType) create(tx Tx, unique bool) (*UsedBaseResourceType, error) {
	var id int
	var savedUnique bool
	err := psql.Insert("base_resource_types").
		Columns("name", "unique_version_history").
		Values(brt.Name, unique).
		Suffix(`
			ON CONFLICT (name) DO UPDATE SET
				name = EXCLUDED.name,
				unique_version_history = EXCLUDED.unique_version_history OR base_resource_types.unique_version_history
			RETURNING id, unique_version_history
		`).
		RunWith(tx).
		QueryRow().
		Scan(&id, &savedUnique)
	if err != nil {
		return nil, err
	}

	baseResourceTypeTable.reloadIfNeeded(tx)

	return &UsedBaseResourceType{ID: id, Name: brt.Name, UniqueVersionHistory: savedUnique}, nil
}
