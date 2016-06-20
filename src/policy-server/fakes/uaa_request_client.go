// This file was generated by counterfeiter
package fakes

import (
	"policy-server/uaa_client"
	"sync"
)

type UAARequestClient struct {
	CheckTokenStub        func(token string) (uaa_client.CheckTokenResponse, error)
	checkTokenMutex       sync.RWMutex
	checkTokenArgsForCall []struct {
		token string
	}
	checkTokenReturns struct {
		result1 uaa_client.CheckTokenResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *UAARequestClient) CheckToken(token string) (uaa_client.CheckTokenResponse, error) {
	fake.checkTokenMutex.Lock()
	fake.checkTokenArgsForCall = append(fake.checkTokenArgsForCall, struct {
		token string
	}{token})
	fake.recordInvocation("CheckToken", []interface{}{token})
	fake.checkTokenMutex.Unlock()
	if fake.CheckTokenStub != nil {
		return fake.CheckTokenStub(token)
	} else {
		return fake.checkTokenReturns.result1, fake.checkTokenReturns.result2
	}
}

func (fake *UAARequestClient) CheckTokenCallCount() int {
	fake.checkTokenMutex.RLock()
	defer fake.checkTokenMutex.RUnlock()
	return len(fake.checkTokenArgsForCall)
}

func (fake *UAARequestClient) CheckTokenArgsForCall(i int) string {
	fake.checkTokenMutex.RLock()
	defer fake.checkTokenMutex.RUnlock()
	return fake.checkTokenArgsForCall[i].token
}

func (fake *UAARequestClient) CheckTokenReturns(result1 uaa_client.CheckTokenResponse, result2 error) {
	fake.CheckTokenStub = nil
	fake.checkTokenReturns = struct {
		result1 uaa_client.CheckTokenResponse
		result2 error
	}{result1, result2}
}

func (fake *UAARequestClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkTokenMutex.RLock()
	defer fake.checkTokenMutex.RUnlock()
	return fake.invocations
}

func (fake *UAARequestClient) recordInvocation(key string, args []interface{}) {
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