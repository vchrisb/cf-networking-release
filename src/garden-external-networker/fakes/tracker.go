// This file was generated by counterfeiter
package fakes

import (
	"garden-external-networker/port_allocator"
	"sync"
)

type Tracker struct {
	AcquireOneStub        func(pool *port_allocator.Pool, handle string) (int, error)
	acquireOneMutex       sync.RWMutex
	acquireOneArgsForCall []struct {
		pool   *port_allocator.Pool
		handle string
	}
	acquireOneReturns struct {
		result1 int
		result2 error
	}
	ReleaseAllStub        func(pool *port_allocator.Pool, handle string) error
	releaseAllMutex       sync.RWMutex
	releaseAllArgsForCall []struct {
		pool   *port_allocator.Pool
		handle string
	}
	releaseAllReturns struct {
		result1 error
	}
	InRangeStub        func(port int) bool
	inRangeMutex       sync.RWMutex
	inRangeArgsForCall []struct {
		port int
	}
	inRangeReturns struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Tracker) AcquireOne(pool *port_allocator.Pool, handle string) (int, error) {
	fake.acquireOneMutex.Lock()
	fake.acquireOneArgsForCall = append(fake.acquireOneArgsForCall, struct {
		pool   *port_allocator.Pool
		handle string
	}{pool, handle})
	fake.recordInvocation("AcquireOne", []interface{}{pool, handle})
	fake.acquireOneMutex.Unlock()
	if fake.AcquireOneStub != nil {
		return fake.AcquireOneStub(pool, handle)
	}
	return fake.acquireOneReturns.result1, fake.acquireOneReturns.result2
}

func (fake *Tracker) AcquireOneCallCount() int {
	fake.acquireOneMutex.RLock()
	defer fake.acquireOneMutex.RUnlock()
	return len(fake.acquireOneArgsForCall)
}

func (fake *Tracker) AcquireOneArgsForCall(i int) (*port_allocator.Pool, string) {
	fake.acquireOneMutex.RLock()
	defer fake.acquireOneMutex.RUnlock()
	return fake.acquireOneArgsForCall[i].pool, fake.acquireOneArgsForCall[i].handle
}

func (fake *Tracker) AcquireOneReturns(result1 int, result2 error) {
	fake.AcquireOneStub = nil
	fake.acquireOneReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *Tracker) ReleaseAll(pool *port_allocator.Pool, handle string) error {
	fake.releaseAllMutex.Lock()
	fake.releaseAllArgsForCall = append(fake.releaseAllArgsForCall, struct {
		pool   *port_allocator.Pool
		handle string
	}{pool, handle})
	fake.recordInvocation("ReleaseAll", []interface{}{pool, handle})
	fake.releaseAllMutex.Unlock()
	if fake.ReleaseAllStub != nil {
		return fake.ReleaseAllStub(pool, handle)
	}
	return fake.releaseAllReturns.result1
}

func (fake *Tracker) ReleaseAllCallCount() int {
	fake.releaseAllMutex.RLock()
	defer fake.releaseAllMutex.RUnlock()
	return len(fake.releaseAllArgsForCall)
}

func (fake *Tracker) ReleaseAllArgsForCall(i int) (*port_allocator.Pool, string) {
	fake.releaseAllMutex.RLock()
	defer fake.releaseAllMutex.RUnlock()
	return fake.releaseAllArgsForCall[i].pool, fake.releaseAllArgsForCall[i].handle
}

func (fake *Tracker) ReleaseAllReturns(result1 error) {
	fake.ReleaseAllStub = nil
	fake.releaseAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *Tracker) InRange(port int) bool {
	fake.inRangeMutex.Lock()
	fake.inRangeArgsForCall = append(fake.inRangeArgsForCall, struct {
		port int
	}{port})
	fake.recordInvocation("InRange", []interface{}{port})
	fake.inRangeMutex.Unlock()
	if fake.InRangeStub != nil {
		return fake.InRangeStub(port)
	}
	return fake.inRangeReturns.result1
}

func (fake *Tracker) InRangeCallCount() int {
	fake.inRangeMutex.RLock()
	defer fake.inRangeMutex.RUnlock()
	return len(fake.inRangeArgsForCall)
}

func (fake *Tracker) InRangeArgsForCall(i int) int {
	fake.inRangeMutex.RLock()
	defer fake.inRangeMutex.RUnlock()
	return fake.inRangeArgsForCall[i].port
}

func (fake *Tracker) InRangeReturns(result1 bool) {
	fake.InRangeStub = nil
	fake.inRangeReturns = struct {
		result1 bool
	}{result1}
}

func (fake *Tracker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireOneMutex.RLock()
	defer fake.acquireOneMutex.RUnlock()
	fake.releaseAllMutex.RLock()
	defer fake.releaseAllMutex.RUnlock()
	fake.inRangeMutex.RLock()
	defer fake.inRangeMutex.RUnlock()
	return fake.invocations
}

func (fake *Tracker) recordInvocation(key string, args []interface{}) {
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
