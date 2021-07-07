// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"
	"time"

	"github.com/concourse/concourse/atc/component"
	"github.com/concourse/concourse/atc/db"
)

type FakeComponent struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	IntervalStub        func() time.Duration
	intervalMutex       sync.RWMutex
	intervalArgsForCall []struct {
	}
	intervalReturns struct {
		result1 time.Duration
	}
	intervalReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	IntervalElapsedStub        func() bool
	intervalElapsedMutex       sync.RWMutex
	intervalElapsedArgsForCall []struct {
	}
	intervalElapsedReturns struct {
		result1 bool
	}
	intervalElapsedReturnsOnCall map[int]struct {
		result1 bool
	}
	LastRanStub        func() time.Time
	lastRanMutex       sync.RWMutex
	lastRanArgsForCall []struct {
	}
	lastRanReturns struct {
		result1 time.Time
	}
	lastRanReturnsOnCall map[int]struct {
		result1 time.Time
	}
	LastRunResultStub        func() string
	lastRunResultMutex       sync.RWMutex
	lastRunResultArgsForCall []struct {
	}
	lastRunResultReturns struct {
		result1 string
	}
	lastRunResultReturnsOnCall map[int]struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	PausedStub        func() bool
	pausedMutex       sync.RWMutex
	pausedArgsForCall []struct {
	}
	pausedReturns struct {
		result1 bool
	}
	pausedReturnsOnCall map[int]struct {
		result1 bool
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct {
	}
	reloadReturns struct {
		result1 bool
		result2 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	UpdateLastRanStub        func(time.Time, component.RunResult) error
	updateLastRanMutex       sync.RWMutex
	updateLastRanArgsForCall []struct {
		arg1 time.Time
		arg2 component.RunResult
	}
	updateLastRanReturns struct {
		result1 error
	}
	updateLastRanReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeComponent) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeComponent) IDCalls(stub func() int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeComponent) IDReturns(result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeComponent) IDReturnsOnCall(i int, result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeComponent) Interval() time.Duration {
	fake.intervalMutex.Lock()
	ret, specificReturn := fake.intervalReturnsOnCall[len(fake.intervalArgsForCall)]
	fake.intervalArgsForCall = append(fake.intervalArgsForCall, struct {
	}{})
	stub := fake.IntervalStub
	fakeReturns := fake.intervalReturns
	fake.recordInvocation("Interval", []interface{}{})
	fake.intervalMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) IntervalCallCount() int {
	fake.intervalMutex.RLock()
	defer fake.intervalMutex.RUnlock()
	return len(fake.intervalArgsForCall)
}

func (fake *FakeComponent) IntervalCalls(stub func() time.Duration) {
	fake.intervalMutex.Lock()
	defer fake.intervalMutex.Unlock()
	fake.IntervalStub = stub
}

func (fake *FakeComponent) IntervalReturns(result1 time.Duration) {
	fake.intervalMutex.Lock()
	defer fake.intervalMutex.Unlock()
	fake.IntervalStub = nil
	fake.intervalReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeComponent) IntervalReturnsOnCall(i int, result1 time.Duration) {
	fake.intervalMutex.Lock()
	defer fake.intervalMutex.Unlock()
	fake.IntervalStub = nil
	if fake.intervalReturnsOnCall == nil {
		fake.intervalReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.intervalReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeComponent) IntervalElapsed() bool {
	fake.intervalElapsedMutex.Lock()
	ret, specificReturn := fake.intervalElapsedReturnsOnCall[len(fake.intervalElapsedArgsForCall)]
	fake.intervalElapsedArgsForCall = append(fake.intervalElapsedArgsForCall, struct {
	}{})
	stub := fake.IntervalElapsedStub
	fakeReturns := fake.intervalElapsedReturns
	fake.recordInvocation("IntervalElapsed", []interface{}{})
	fake.intervalElapsedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) IntervalElapsedCallCount() int {
	fake.intervalElapsedMutex.RLock()
	defer fake.intervalElapsedMutex.RUnlock()
	return len(fake.intervalElapsedArgsForCall)
}

func (fake *FakeComponent) IntervalElapsedCalls(stub func() bool) {
	fake.intervalElapsedMutex.Lock()
	defer fake.intervalElapsedMutex.Unlock()
	fake.IntervalElapsedStub = stub
}

func (fake *FakeComponent) IntervalElapsedReturns(result1 bool) {
	fake.intervalElapsedMutex.Lock()
	defer fake.intervalElapsedMutex.Unlock()
	fake.IntervalElapsedStub = nil
	fake.intervalElapsedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeComponent) IntervalElapsedReturnsOnCall(i int, result1 bool) {
	fake.intervalElapsedMutex.Lock()
	defer fake.intervalElapsedMutex.Unlock()
	fake.IntervalElapsedStub = nil
	if fake.intervalElapsedReturnsOnCall == nil {
		fake.intervalElapsedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.intervalElapsedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeComponent) LastRan() time.Time {
	fake.lastRanMutex.Lock()
	ret, specificReturn := fake.lastRanReturnsOnCall[len(fake.lastRanArgsForCall)]
	fake.lastRanArgsForCall = append(fake.lastRanArgsForCall, struct {
	}{})
	stub := fake.LastRanStub
	fakeReturns := fake.lastRanReturns
	fake.recordInvocation("LastRan", []interface{}{})
	fake.lastRanMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) LastRanCallCount() int {
	fake.lastRanMutex.RLock()
	defer fake.lastRanMutex.RUnlock()
	return len(fake.lastRanArgsForCall)
}

func (fake *FakeComponent) LastRanCalls(stub func() time.Time) {
	fake.lastRanMutex.Lock()
	defer fake.lastRanMutex.Unlock()
	fake.LastRanStub = stub
}

func (fake *FakeComponent) LastRanReturns(result1 time.Time) {
	fake.lastRanMutex.Lock()
	defer fake.lastRanMutex.Unlock()
	fake.LastRanStub = nil
	fake.lastRanReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeComponent) LastRanReturnsOnCall(i int, result1 time.Time) {
	fake.lastRanMutex.Lock()
	defer fake.lastRanMutex.Unlock()
	fake.LastRanStub = nil
	if fake.lastRanReturnsOnCall == nil {
		fake.lastRanReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.lastRanReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeComponent) LastRunResult() string {
	fake.lastRunResultMutex.Lock()
	ret, specificReturn := fake.lastRunResultReturnsOnCall[len(fake.lastRunResultArgsForCall)]
	fake.lastRunResultArgsForCall = append(fake.lastRunResultArgsForCall, struct {
	}{})
	stub := fake.LastRunResultStub
	fakeReturns := fake.lastRunResultReturns
	fake.recordInvocation("LastRunResult", []interface{}{})
	fake.lastRunResultMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) LastRunResultCallCount() int {
	fake.lastRunResultMutex.RLock()
	defer fake.lastRunResultMutex.RUnlock()
	return len(fake.lastRunResultArgsForCall)
}

func (fake *FakeComponent) LastRunResultCalls(stub func() string) {
	fake.lastRunResultMutex.Lock()
	defer fake.lastRunResultMutex.Unlock()
	fake.LastRunResultStub = stub
}

func (fake *FakeComponent) LastRunResultReturns(result1 string) {
	fake.lastRunResultMutex.Lock()
	defer fake.lastRunResultMutex.Unlock()
	fake.LastRunResultStub = nil
	fake.lastRunResultReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) LastRunResultReturnsOnCall(i int, result1 string) {
	fake.lastRunResultMutex.Lock()
	defer fake.lastRunResultMutex.Unlock()
	fake.LastRunResultStub = nil
	if fake.lastRunResultReturnsOnCall == nil {
		fake.lastRunResultReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.lastRunResultReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeComponent) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeComponent) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeComponent) Paused() bool {
	fake.pausedMutex.Lock()
	ret, specificReturn := fake.pausedReturnsOnCall[len(fake.pausedArgsForCall)]
	fake.pausedArgsForCall = append(fake.pausedArgsForCall, struct {
	}{})
	stub := fake.PausedStub
	fakeReturns := fake.pausedReturns
	fake.recordInvocation("Paused", []interface{}{})
	fake.pausedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) PausedCallCount() int {
	fake.pausedMutex.RLock()
	defer fake.pausedMutex.RUnlock()
	return len(fake.pausedArgsForCall)
}

func (fake *FakeComponent) PausedCalls(stub func() bool) {
	fake.pausedMutex.Lock()
	defer fake.pausedMutex.Unlock()
	fake.PausedStub = stub
}

func (fake *FakeComponent) PausedReturns(result1 bool) {
	fake.pausedMutex.Lock()
	defer fake.pausedMutex.Unlock()
	fake.PausedStub = nil
	fake.pausedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeComponent) PausedReturnsOnCall(i int, result1 bool) {
	fake.pausedMutex.Lock()
	defer fake.pausedMutex.Unlock()
	fake.PausedStub = nil
	if fake.pausedReturnsOnCall == nil {
		fake.pausedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.pausedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeComponent) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct {
	}{})
	stub := fake.ReloadStub
	fakeReturns := fake.reloadReturns
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeComponent) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeComponent) ReloadCalls(stub func() (bool, error)) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = stub
}

func (fake *FakeComponent) ReloadReturns(result1 bool, result2 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeComponent) ReloadReturnsOnCall(i int, result1 bool, result2 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeComponent) UpdateLastRan(arg1 time.Time, arg2 component.RunResult) error {
	fake.updateLastRanMutex.Lock()
	ret, specificReturn := fake.updateLastRanReturnsOnCall[len(fake.updateLastRanArgsForCall)]
	fake.updateLastRanArgsForCall = append(fake.updateLastRanArgsForCall, struct {
		arg1 time.Time
		arg2 component.RunResult
	}{arg1, arg2})
	stub := fake.UpdateLastRanStub
	fakeReturns := fake.updateLastRanReturns
	fake.recordInvocation("UpdateLastRan", []interface{}{arg1, arg2})
	fake.updateLastRanMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeComponent) UpdateLastRanCallCount() int {
	fake.updateLastRanMutex.RLock()
	defer fake.updateLastRanMutex.RUnlock()
	return len(fake.updateLastRanArgsForCall)
}

func (fake *FakeComponent) UpdateLastRanCalls(stub func(time.Time, component.RunResult) error) {
	fake.updateLastRanMutex.Lock()
	defer fake.updateLastRanMutex.Unlock()
	fake.UpdateLastRanStub = stub
}

func (fake *FakeComponent) UpdateLastRanArgsForCall(i int) (time.Time, component.RunResult) {
	fake.updateLastRanMutex.RLock()
	defer fake.updateLastRanMutex.RUnlock()
	argsForCall := fake.updateLastRanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeComponent) UpdateLastRanReturns(result1 error) {
	fake.updateLastRanMutex.Lock()
	defer fake.updateLastRanMutex.Unlock()
	fake.UpdateLastRanStub = nil
	fake.updateLastRanReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeComponent) UpdateLastRanReturnsOnCall(i int, result1 error) {
	fake.updateLastRanMutex.Lock()
	defer fake.updateLastRanMutex.Unlock()
	fake.UpdateLastRanStub = nil
	if fake.updateLastRanReturnsOnCall == nil {
		fake.updateLastRanReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateLastRanReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeComponent) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.intervalMutex.RLock()
	defer fake.intervalMutex.RUnlock()
	fake.intervalElapsedMutex.RLock()
	defer fake.intervalElapsedMutex.RUnlock()
	fake.lastRanMutex.RLock()
	defer fake.lastRanMutex.RUnlock()
	fake.lastRunResultMutex.RLock()
	defer fake.lastRunResultMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.pausedMutex.RLock()
	defer fake.pausedMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.updateLastRanMutex.RLock()
	defer fake.updateLastRanMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeComponent) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.Component = new(FakeComponent)
