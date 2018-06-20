// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"lib/policy_client"
	"sync"
)

type PolicyClient struct {
	GetPoliciesByIDStub        func(ids ...string) ([]policy_client.Policy, error)
	getPoliciesByIDMutex       sync.RWMutex
	getPoliciesByIDArgsForCall []struct {
		ids []string
	}
	getPoliciesByIDReturns struct {
		result1 []policy_client.Policy
		result2 error
	}
	getPoliciesByIDReturnsOnCall map[int]struct {
		result1 []policy_client.Policy
		result2 error
	}
	CreateOrGetTagStub        func(id, groupType string) (string, error)
	createOrGetTagMutex       sync.RWMutex
	createOrGetTagArgsForCall []struct {
		id        string
		groupType string
	}
	createOrGetTagReturns struct {
		result1 string
		result2 error
	}
	createOrGetTagReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyClient) GetPoliciesByID(ids ...string) ([]policy_client.Policy, error) {
	fake.getPoliciesByIDMutex.Lock()
	ret, specificReturn := fake.getPoliciesByIDReturnsOnCall[len(fake.getPoliciesByIDArgsForCall)]
	fake.getPoliciesByIDArgsForCall = append(fake.getPoliciesByIDArgsForCall, struct {
		ids []string
	}{ids})
	fake.recordInvocation("GetPoliciesByID", []interface{}{ids})
	fake.getPoliciesByIDMutex.Unlock()
	if fake.GetPoliciesByIDStub != nil {
		return fake.GetPoliciesByIDStub(ids...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getPoliciesByIDReturns.result1, fake.getPoliciesByIDReturns.result2
}

func (fake *PolicyClient) GetPoliciesByIDCallCount() int {
	fake.getPoliciesByIDMutex.RLock()
	defer fake.getPoliciesByIDMutex.RUnlock()
	return len(fake.getPoliciesByIDArgsForCall)
}

func (fake *PolicyClient) GetPoliciesByIDArgsForCall(i int) []string {
	fake.getPoliciesByIDMutex.RLock()
	defer fake.getPoliciesByIDMutex.RUnlock()
	return fake.getPoliciesByIDArgsForCall[i].ids
}

func (fake *PolicyClient) GetPoliciesByIDReturns(result1 []policy_client.Policy, result2 error) {
	fake.GetPoliciesByIDStub = nil
	fake.getPoliciesByIDReturns = struct {
		result1 []policy_client.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) GetPoliciesByIDReturnsOnCall(i int, result1 []policy_client.Policy, result2 error) {
	fake.GetPoliciesByIDStub = nil
	if fake.getPoliciesByIDReturnsOnCall == nil {
		fake.getPoliciesByIDReturnsOnCall = make(map[int]struct {
			result1 []policy_client.Policy
			result2 error
		})
	}
	fake.getPoliciesByIDReturnsOnCall[i] = struct {
		result1 []policy_client.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) CreateOrGetTag(id string, groupType string) (string, error) {
	fake.createOrGetTagMutex.Lock()
	ret, specificReturn := fake.createOrGetTagReturnsOnCall[len(fake.createOrGetTagArgsForCall)]
	fake.createOrGetTagArgsForCall = append(fake.createOrGetTagArgsForCall, struct {
		id        string
		groupType string
	}{id, groupType})
	fake.recordInvocation("CreateOrGetTag", []interface{}{id, groupType})
	fake.createOrGetTagMutex.Unlock()
	if fake.CreateOrGetTagStub != nil {
		return fake.CreateOrGetTagStub(id, groupType)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createOrGetTagReturns.result1, fake.createOrGetTagReturns.result2
}

func (fake *PolicyClient) CreateOrGetTagCallCount() int {
	fake.createOrGetTagMutex.RLock()
	defer fake.createOrGetTagMutex.RUnlock()
	return len(fake.createOrGetTagArgsForCall)
}

func (fake *PolicyClient) CreateOrGetTagArgsForCall(i int) (string, string) {
	fake.createOrGetTagMutex.RLock()
	defer fake.createOrGetTagMutex.RUnlock()
	return fake.createOrGetTagArgsForCall[i].id, fake.createOrGetTagArgsForCall[i].groupType
}

func (fake *PolicyClient) CreateOrGetTagReturns(result1 string, result2 error) {
	fake.CreateOrGetTagStub = nil
	fake.createOrGetTagReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) CreateOrGetTagReturnsOnCall(i int, result1 string, result2 error) {
	fake.CreateOrGetTagStub = nil
	if fake.createOrGetTagReturnsOnCall == nil {
		fake.createOrGetTagReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createOrGetTagReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPoliciesByIDMutex.RLock()
	defer fake.getPoliciesByIDMutex.RUnlock()
	fake.createOrGetTagMutex.RLock()
	defer fake.createOrGetTagMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyClient) recordInvocation(key string, args []interface{}) {
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
