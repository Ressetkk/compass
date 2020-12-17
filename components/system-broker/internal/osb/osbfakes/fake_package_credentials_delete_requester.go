// Code generated by counterfeiter. DO NOT EDIT.
package osbfakes

import (
	"context"
	"sync"

	"github.com/kyma-incubator/compass/components/system-broker/internal/director"
)

type FakePackageCredentialsDeleteRequester struct {
	RequestPackageInstanceCredentialsDeletionStub        func(context.Context, *director.PackageInstanceAuthDeletionInput) (*director.PackageInstanceAuthDeletionOutput, error)
	requestPackageInstanceCredentialsDeletionMutex       sync.RWMutex
	requestPackageInstanceCredentialsDeletionArgsForCall []struct {
		arg1 context.Context
		arg2 *director.PackageInstanceAuthDeletionInput
	}
	requestPackageInstanceCredentialsDeletionReturns struct {
		result1 *director.PackageInstanceAuthDeletionOutput
		result2 error
	}
	requestPackageInstanceCredentialsDeletionReturnsOnCall map[int]struct {
		result1 *director.PackageInstanceAuthDeletionOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePackageCredentialsDeleteRequester) RequestPackageInstanceCredentialsDeletion(arg1 context.Context, arg2 *director.PackageInstanceAuthDeletionInput) (*director.PackageInstanceAuthDeletionOutput, error) {
	fake.requestPackageInstanceCredentialsDeletionMutex.Lock()
	ret, specificReturn := fake.requestPackageInstanceCredentialsDeletionReturnsOnCall[len(fake.requestPackageInstanceCredentialsDeletionArgsForCall)]
	fake.requestPackageInstanceCredentialsDeletionArgsForCall = append(fake.requestPackageInstanceCredentialsDeletionArgsForCall, struct {
		arg1 context.Context
		arg2 *director.PackageInstanceAuthDeletionInput
	}{arg1, arg2})
	fake.recordInvocation("RequestPackageInstanceCredentialsDeletion", []interface{}{arg1, arg2})
	fake.requestPackageInstanceCredentialsDeletionMutex.Unlock()
	if fake.RequestPackageInstanceCredentialsDeletionStub != nil {
		return fake.RequestPackageInstanceCredentialsDeletionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requestPackageInstanceCredentialsDeletionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePackageCredentialsDeleteRequester) RequestPackageInstanceCredentialsDeletionCallCount() int {
	fake.requestPackageInstanceCredentialsDeletionMutex.RLock()
	defer fake.requestPackageInstanceCredentialsDeletionMutex.RUnlock()
	return len(fake.requestPackageInstanceCredentialsDeletionArgsForCall)
}

func (fake *FakePackageCredentialsDeleteRequester) RequestPackageInstanceCredentialsDeletionCalls(stub func(context.Context, *director.PackageInstanceAuthDeletionInput) (*director.PackageInstanceAuthDeletionOutput, error)) {
	fake.requestPackageInstanceCredentialsDeletionMutex.Lock()
	defer fake.requestPackageInstanceCredentialsDeletionMutex.Unlock()
	fake.RequestPackageInstanceCredentialsDeletionStub = stub
}

func (fake *FakePackageCredentialsDeleteRequester) RequestPackageInstanceCredentialsDeletionArgsForCall(i int) (context.Context, *director.PackageInstanceAuthDeletionInput) {
	fake.requestPackageInstanceCredentialsDeletionMutex.RLock()
	defer fake.requestPackageInstanceCredentialsDeletionMutex.RUnlock()
	argsForCall := fake.requestPackageInstanceCredentialsDeletionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePackageCredentialsDeleteRequester) RequestPackageInstanceCredentialsDeletionReturns(result1 *director.PackageInstanceAuthDeletionOutput, result2 error) {
	fake.requestPackageInstanceCredentialsDeletionMutex.Lock()
	defer fake.requestPackageInstanceCredentialsDeletionMutex.Unlock()
	fake.RequestPackageInstanceCredentialsDeletionStub = nil
	fake.requestPackageInstanceCredentialsDeletionReturns = struct {
		result1 *director.PackageInstanceAuthDeletionOutput
		result2 error
	}{result1, result2}
}

func (fake *FakePackageCredentialsDeleteRequester) RequestPackageInstanceCredentialsDeletionReturnsOnCall(i int, result1 *director.PackageInstanceAuthDeletionOutput, result2 error) {
	fake.requestPackageInstanceCredentialsDeletionMutex.Lock()
	defer fake.requestPackageInstanceCredentialsDeletionMutex.Unlock()
	fake.RequestPackageInstanceCredentialsDeletionStub = nil
	if fake.requestPackageInstanceCredentialsDeletionReturnsOnCall == nil {
		fake.requestPackageInstanceCredentialsDeletionReturnsOnCall = make(map[int]struct {
			result1 *director.PackageInstanceAuthDeletionOutput
			result2 error
		})
	}
	fake.requestPackageInstanceCredentialsDeletionReturnsOnCall[i] = struct {
		result1 *director.PackageInstanceAuthDeletionOutput
		result2 error
	}{result1, result2}
}

func (fake *FakePackageCredentialsDeleteRequester) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.requestPackageInstanceCredentialsDeletionMutex.RLock()
	defer fake.requestPackageInstanceCredentialsDeletionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePackageCredentialsDeleteRequester) recordInvocation(key string, args []interface{}) {
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
