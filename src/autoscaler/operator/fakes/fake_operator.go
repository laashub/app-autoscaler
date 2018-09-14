// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"autoscaler/operator"
	"sync"
)

type FakeOperator struct {
	OperateStub        func()
	operateMutex       sync.RWMutex
	operateArgsForCall []struct{}
	invocations        map[string][][]interface{}
	invocationsMutex   sync.RWMutex
}

func (fake *FakeOperator) Operate() {
	fake.operateMutex.Lock()
	fake.operateArgsForCall = append(fake.operateArgsForCall, struct{}{})
	fake.recordInvocation("Operate", []interface{}{})
	fake.operateMutex.Unlock()
	if fake.OperateStub != nil {
		fake.OperateStub()
	}
}

func (fake *FakeOperator) OperateCallCount() int {
	fake.operateMutex.RLock()
	defer fake.operateMutex.RUnlock()
	return len(fake.operateArgsForCall)
}

func (fake *FakeOperator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.operateMutex.RLock()
	defer fake.operateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOperator) recordInvocation(key string, args []interface{}) {
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

var _ operator.Operator = new(FakeOperator)