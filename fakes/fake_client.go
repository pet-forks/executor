// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry-incubator/executor"
)

type FakeClient struct {
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	AllocateContainersStub        func(requests []executor.Container) (map[string]string, error)
	allocateContainersMutex       sync.RWMutex
	allocateContainersArgsForCall []struct {
		requests []executor.Container
	}
	allocateContainersReturns struct {
		result1 map[string]string
		result2 error
	}
	GetContainerStub        func(guid string) (executor.Container, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		guid string
	}
	getContainerReturns struct {
		result1 executor.Container
		result2 error
	}
	RunContainerStub        func(guid string) error
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		guid string
	}
	runContainerReturns struct {
		result1 error
	}
	StopContainerStub        func(guid string) error
	stopContainerMutex       sync.RWMutex
	stopContainerArgsForCall []struct {
		guid string
	}
	stopContainerReturns struct {
		result1 error
	}
	DeleteContainerStub        func(guid string) error
	deleteContainerMutex       sync.RWMutex
	deleteContainerArgsForCall []struct {
		guid string
	}
	deleteContainerReturns struct {
		result1 error
	}
	ListContainersStub        func(executor.Tags) ([]executor.Container, error)
	listContainersMutex       sync.RWMutex
	listContainersArgsForCall []struct {
		arg1 executor.Tags
	}
	listContainersReturns struct {
		result1 []executor.Container
		result2 error
	}
	RemainingResourcesStub        func() (executor.ExecutorResources, error)
	remainingResourcesMutex       sync.RWMutex
	remainingResourcesArgsForCall []struct{}
	remainingResourcesReturns     struct {
		result1 executor.ExecutorResources
		result2 error
	}
	TotalResourcesStub        func() (executor.ExecutorResources, error)
	totalResourcesMutex       sync.RWMutex
	totalResourcesArgsForCall []struct{}
	totalResourcesReturns     struct {
		result1 executor.ExecutorResources
		result2 error
	}
	GetFilesStub        func(guid string, path string) (io.ReadCloser, error)
	getFilesMutex       sync.RWMutex
	getFilesArgsForCall []struct {
		guid string
		path string
	}
	getFilesReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	SubscribeToEventsStub        func() (<-chan executor.Event, error)
	subscribeToEventsMutex       sync.RWMutex
	subscribeToEventsArgsForCall []struct{}
	subscribeToEventsReturns     struct {
		result1 <-chan executor.Event
		result2 error
	}
}

func (fake *FakeClient) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeClient) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AllocateContainers(requests []executor.Container) (map[string]string, error) {
	fake.allocateContainersMutex.Lock()
	fake.allocateContainersArgsForCall = append(fake.allocateContainersArgsForCall, struct {
		requests []executor.Container
	}{requests})
	fake.allocateContainersMutex.Unlock()
	if fake.AllocateContainersStub != nil {
		return fake.AllocateContainersStub(requests)
	} else {
		return fake.allocateContainersReturns.result1, fake.allocateContainersReturns.result2
	}
}

func (fake *FakeClient) AllocateContainersCallCount() int {
	fake.allocateContainersMutex.RLock()
	defer fake.allocateContainersMutex.RUnlock()
	return len(fake.allocateContainersArgsForCall)
}

func (fake *FakeClient) AllocateContainersArgsForCall(i int) []executor.Container {
	fake.allocateContainersMutex.RLock()
	defer fake.allocateContainersMutex.RUnlock()
	return fake.allocateContainersArgsForCall[i].requests
}

func (fake *FakeClient) AllocateContainersReturns(result1 map[string]string, result2 error) {
	fake.AllocateContainersStub = nil
	fake.allocateContainersReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetContainer(guid string) (executor.Container, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		guid string
	}{guid})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(guid)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2
	}
}

func (fake *FakeClient) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeClient) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].guid
}

func (fake *FakeClient) GetContainerReturns(result1 executor.Container, result2 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RunContainer(guid string) error {
	fake.runContainerMutex.Lock()
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		guid string
	}{guid})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(guid)
	} else {
		return fake.runContainerReturns.result1
	}
}

func (fake *FakeClient) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeClient) RunContainerArgsForCall(i int) string {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return fake.runContainerArgsForCall[i].guid
}

func (fake *FakeClient) RunContainerReturns(result1 error) {
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopContainer(guid string) error {
	fake.stopContainerMutex.Lock()
	fake.stopContainerArgsForCall = append(fake.stopContainerArgsForCall, struct {
		guid string
	}{guid})
	fake.stopContainerMutex.Unlock()
	if fake.StopContainerStub != nil {
		return fake.StopContainerStub(guid)
	} else {
		return fake.stopContainerReturns.result1
	}
}

func (fake *FakeClient) StopContainerCallCount() int {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return len(fake.stopContainerArgsForCall)
}

func (fake *FakeClient) StopContainerArgsForCall(i int) string {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return fake.stopContainerArgsForCall[i].guid
}

func (fake *FakeClient) StopContainerReturns(result1 error) {
	fake.StopContainerStub = nil
	fake.stopContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeleteContainer(guid string) error {
	fake.deleteContainerMutex.Lock()
	fake.deleteContainerArgsForCall = append(fake.deleteContainerArgsForCall, struct {
		guid string
	}{guid})
	fake.deleteContainerMutex.Unlock()
	if fake.DeleteContainerStub != nil {
		return fake.DeleteContainerStub(guid)
	} else {
		return fake.deleteContainerReturns.result1
	}
}

func (fake *FakeClient) DeleteContainerCallCount() int {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return len(fake.deleteContainerArgsForCall)
}

func (fake *FakeClient) DeleteContainerArgsForCall(i int) string {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return fake.deleteContainerArgsForCall[i].guid
}

func (fake *FakeClient) DeleteContainerReturns(result1 error) {
	fake.DeleteContainerStub = nil
	fake.deleteContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ListContainers(arg1 executor.Tags) ([]executor.Container, error) {
	fake.listContainersMutex.Lock()
	fake.listContainersArgsForCall = append(fake.listContainersArgsForCall, struct {
		arg1 executor.Tags
	}{arg1})
	fake.listContainersMutex.Unlock()
	if fake.ListContainersStub != nil {
		return fake.ListContainersStub(arg1)
	} else {
		return fake.listContainersReturns.result1, fake.listContainersReturns.result2
	}
}

func (fake *FakeClient) ListContainersCallCount() int {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return len(fake.listContainersArgsForCall)
}

func (fake *FakeClient) ListContainersArgsForCall(i int) executor.Tags {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return fake.listContainersArgsForCall[i].arg1
}

func (fake *FakeClient) ListContainersReturns(result1 []executor.Container, result2 error) {
	fake.ListContainersStub = nil
	fake.listContainersReturns = struct {
		result1 []executor.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RemainingResources() (executor.ExecutorResources, error) {
	fake.remainingResourcesMutex.Lock()
	fake.remainingResourcesArgsForCall = append(fake.remainingResourcesArgsForCall, struct{}{})
	fake.remainingResourcesMutex.Unlock()
	if fake.RemainingResourcesStub != nil {
		return fake.RemainingResourcesStub()
	} else {
		return fake.remainingResourcesReturns.result1, fake.remainingResourcesReturns.result2
	}
}

func (fake *FakeClient) RemainingResourcesCallCount() int {
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	return len(fake.remainingResourcesArgsForCall)
}

func (fake *FakeClient) RemainingResourcesReturns(result1 executor.ExecutorResources, result2 error) {
	fake.RemainingResourcesStub = nil
	fake.remainingResourcesReturns = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TotalResources() (executor.ExecutorResources, error) {
	fake.totalResourcesMutex.Lock()
	fake.totalResourcesArgsForCall = append(fake.totalResourcesArgsForCall, struct{}{})
	fake.totalResourcesMutex.Unlock()
	if fake.TotalResourcesStub != nil {
		return fake.TotalResourcesStub()
	} else {
		return fake.totalResourcesReturns.result1, fake.totalResourcesReturns.result2
	}
}

func (fake *FakeClient) TotalResourcesCallCount() int {
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	return len(fake.totalResourcesArgsForCall)
}

func (fake *FakeClient) TotalResourcesReturns(result1 executor.ExecutorResources, result2 error) {
	fake.TotalResourcesStub = nil
	fake.totalResourcesReturns = struct {
		result1 executor.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetFiles(guid string, path string) (io.ReadCloser, error) {
	fake.getFilesMutex.Lock()
	fake.getFilesArgsForCall = append(fake.getFilesArgsForCall, struct {
		guid string
		path string
	}{guid, path})
	fake.getFilesMutex.Unlock()
	if fake.GetFilesStub != nil {
		return fake.GetFilesStub(guid, path)
	} else {
		return fake.getFilesReturns.result1, fake.getFilesReturns.result2
	}
}

func (fake *FakeClient) GetFilesCallCount() int {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	return len(fake.getFilesArgsForCall)
}

func (fake *FakeClient) GetFilesArgsForCall(i int) (string, string) {
	fake.getFilesMutex.RLock()
	defer fake.getFilesMutex.RUnlock()
	return fake.getFilesArgsForCall[i].guid, fake.getFilesArgsForCall[i].path
}

func (fake *FakeClient) GetFilesReturns(result1 io.ReadCloser, result2 error) {
	fake.GetFilesStub = nil
	fake.getFilesReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SubscribeToEvents() (<-chan executor.Event, error) {
	fake.subscribeToEventsMutex.Lock()
	fake.subscribeToEventsArgsForCall = append(fake.subscribeToEventsArgsForCall, struct{}{})
	fake.subscribeToEventsMutex.Unlock()
	if fake.SubscribeToEventsStub != nil {
		return fake.SubscribeToEventsStub()
	} else {
		return fake.subscribeToEventsReturns.result1, fake.subscribeToEventsReturns.result2
	}
}

func (fake *FakeClient) SubscribeToEventsCallCount() int {
	fake.subscribeToEventsMutex.RLock()
	defer fake.subscribeToEventsMutex.RUnlock()
	return len(fake.subscribeToEventsArgsForCall)
}

func (fake *FakeClient) SubscribeToEventsReturns(result1 <-chan executor.Event, result2 error) {
	fake.SubscribeToEventsStub = nil
	fake.subscribeToEventsReturns = struct {
		result1 <-chan executor.Event
		result2 error
	}{result1, result2}
}

var _ executor.Client = new(FakeClient)
