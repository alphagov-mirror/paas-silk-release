// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type Locker struct {
	LockStub        func() error
	lockMutex       sync.RWMutex
	lockArgsForCall []struct{}
	lockReturns     struct {
		result1 error
	}
	lockReturnsOnCall map[int]struct {
		result1 error
	}
	UnlockStub        func() error
	unlockMutex       sync.RWMutex
	unlockArgsForCall []struct{}
	unlockReturns     struct {
		result1 error
	}
	unlockReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Locker) Lock() error {
	fake.lockMutex.Lock()
	ret, specificReturn := fake.lockReturnsOnCall[len(fake.lockArgsForCall)]
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct{}{})
	fake.recordInvocation("Lock", []interface{}{})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.lockReturns.result1
}

func (fake *Locker) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *Locker) LockReturns(result1 error) {
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 error
	}{result1}
}

func (fake *Locker) LockReturnsOnCall(i int, result1 error) {
	fake.LockStub = nil
	if fake.lockReturnsOnCall == nil {
		fake.lockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.lockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Locker) Unlock() error {
	fake.unlockMutex.Lock()
	ret, specificReturn := fake.unlockReturnsOnCall[len(fake.unlockArgsForCall)]
	fake.unlockArgsForCall = append(fake.unlockArgsForCall, struct{}{})
	fake.recordInvocation("Unlock", []interface{}{})
	fake.unlockMutex.Unlock()
	if fake.UnlockStub != nil {
		return fake.UnlockStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unlockReturns.result1
}

func (fake *Locker) UnlockCallCount() int {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return len(fake.unlockArgsForCall)
}

func (fake *Locker) UnlockReturns(result1 error) {
	fake.UnlockStub = nil
	fake.unlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *Locker) UnlockReturnsOnCall(i int, result1 error) {
	fake.UnlockStub = nil
	if fake.unlockReturnsOnCall == nil {
		fake.unlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Locker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return fake.invocations
}

func (fake *Locker) recordInvocation(key string, args []interface{}) {
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