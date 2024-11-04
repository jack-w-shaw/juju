// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/storage (interfaces: ProviderRegistry)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination internal_storage_mock_test.go github.com/juju/juju/internal/storage ProviderRegistry
//

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	storage "github.com/juju/juju/internal/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockProviderRegistry is a mock of ProviderRegistry interface.
type MockProviderRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockProviderRegistryMockRecorder
}

// MockProviderRegistryMockRecorder is the mock recorder for MockProviderRegistry.
type MockProviderRegistryMockRecorder struct {
	mock *MockProviderRegistry
}

// NewMockProviderRegistry creates a new mock instance.
func NewMockProviderRegistry(ctrl *gomock.Controller) *MockProviderRegistry {
	mock := &MockProviderRegistry{ctrl: ctrl}
	mock.recorder = &MockProviderRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderRegistry) EXPECT() *MockProviderRegistryMockRecorder {
	return m.recorder
}

// StorageProvider mocks base method.
func (m *MockProviderRegistry) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockProviderRegistryMockRecorder) StorageProvider(arg0 any) *MockProviderRegistryStorageProviderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockProviderRegistry)(nil).StorageProvider), arg0)
	return &MockProviderRegistryStorageProviderCall{Call: call}
}

// MockProviderRegistryStorageProviderCall wrap *gomock.Call
type MockProviderRegistryStorageProviderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderRegistryStorageProviderCall) Return(arg0 storage.Provider, arg1 error) *MockProviderRegistryStorageProviderCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderRegistryStorageProviderCall) Do(f func(storage.ProviderType) (storage.Provider, error)) *MockProviderRegistryStorageProviderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderRegistryStorageProviderCall) DoAndReturn(f func(storage.ProviderType) (storage.Provider, error)) *MockProviderRegistryStorageProviderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StorageProviderTypes mocks base method.
func (m *MockProviderRegistry) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockProviderRegistryMockRecorder) StorageProviderTypes() *MockProviderRegistryStorageProviderTypesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockProviderRegistry)(nil).StorageProviderTypes))
	return &MockProviderRegistryStorageProviderTypesCall{Call: call}
}

// MockProviderRegistryStorageProviderTypesCall wrap *gomock.Call
type MockProviderRegistryStorageProviderTypesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderRegistryStorageProviderTypesCall) Return(arg0 []storage.ProviderType, arg1 error) *MockProviderRegistryStorageProviderTypesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderRegistryStorageProviderTypesCall) Do(f func() ([]storage.ProviderType, error)) *MockProviderRegistryStorageProviderTypesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderRegistryStorageProviderTypesCall) DoAndReturn(f func() ([]storage.ProviderType, error)) *MockProviderRegistryStorageProviderTypesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
