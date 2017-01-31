// This file was generated by counterfeiter
package fakes

import (
	"lib/models"
	"sync"
)

type PolicyClient struct {
	GetPoliciesStub        func() ([]models.Policy, error)
	getPoliciesMutex       sync.RWMutex
	getPoliciesArgsForCall []struct{}
	getPoliciesReturns     struct {
		result1 []models.Policy
		result2 error
	}
	GetPoliciesByIDStub        func(ids ...string) ([]models.Policy, error)
	getPoliciesByIDMutex       sync.RWMutex
	getPoliciesByIDArgsForCall []struct {
		ids []string
	}
	getPoliciesByIDReturns struct {
		result1 []models.Policy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyClient) GetPolicies() ([]models.Policy, error) {
	fake.getPoliciesMutex.Lock()
	fake.getPoliciesArgsForCall = append(fake.getPoliciesArgsForCall, struct{}{})
	fake.recordInvocation("GetPolicies", []interface{}{})
	fake.getPoliciesMutex.Unlock()
	if fake.GetPoliciesStub != nil {
		return fake.GetPoliciesStub()
	} else {
		return fake.getPoliciesReturns.result1, fake.getPoliciesReturns.result2
	}
}

func (fake *PolicyClient) GetPoliciesCallCount() int {
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	return len(fake.getPoliciesArgsForCall)
}

func (fake *PolicyClient) GetPoliciesReturns(result1 []models.Policy, result2 error) {
	fake.GetPoliciesStub = nil
	fake.getPoliciesReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) GetPoliciesByID(ids ...string) ([]models.Policy, error) {
	fake.getPoliciesByIDMutex.Lock()
	fake.getPoliciesByIDArgsForCall = append(fake.getPoliciesByIDArgsForCall, struct {
		ids []string
	}{ids})
	fake.recordInvocation("GetPoliciesByID", []interface{}{ids})
	fake.getPoliciesByIDMutex.Unlock()
	if fake.GetPoliciesByIDStub != nil {
		return fake.GetPoliciesByIDStub(ids...)
	} else {
		return fake.getPoliciesByIDReturns.result1, fake.getPoliciesByIDReturns.result2
	}
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

func (fake *PolicyClient) GetPoliciesByIDReturns(result1 []models.Policy, result2 error) {
	fake.GetPoliciesByIDStub = nil
	fake.getPoliciesByIDReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPoliciesMutex.RLock()
	defer fake.getPoliciesMutex.RUnlock()
	fake.getPoliciesByIDMutex.RLock()
	defer fake.getPoliciesByIDMutex.RUnlock()
	return fake.invocations
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
