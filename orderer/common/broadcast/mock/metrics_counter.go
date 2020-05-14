// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric-gm/common/metrics"
)

type MetricsCounter struct {
	WithStub        func(labelValues ...string) metrics.Counter
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		labelValues []string
	}
	withReturns struct {
		result1 metrics.Counter
	}
	withReturnsOnCall map[int]struct {
		result1 metrics.Counter
	}
	AddStub        func(delta float64)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		delta float64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetricsCounter) With(labelValues ...string) metrics.Counter {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct {
		labelValues []string
	}{labelValues})
	fake.recordInvocation("With", []interface{}{labelValues})
	fake.withMutex.Unlock()
	if fake.WithStub != nil {
		return fake.WithStub(labelValues...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.withReturns.result1
}

func (fake *MetricsCounter) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

func (fake *MetricsCounter) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return fake.withArgsForCall[i].labelValues
}

func (fake *MetricsCounter) WithReturns(result1 metrics.Counter) {
	fake.WithStub = nil
	fake.withReturns = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *MetricsCounter) WithReturnsOnCall(i int, result1 metrics.Counter) {
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct {
			result1 metrics.Counter
		})
	}
	fake.withReturnsOnCall[i] = struct {
		result1 metrics.Counter
	}{result1}
}

func (fake *MetricsCounter) Add(delta float64) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		delta float64
	}{delta})
	fake.recordInvocation("Add", []interface{}{delta})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(delta)
	}
}

func (fake *MetricsCounter) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *MetricsCounter) AddArgsForCall(i int) float64 {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].delta
}

func (fake *MetricsCounter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetricsCounter) recordInvocation(key string, args []interface{}) {
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
