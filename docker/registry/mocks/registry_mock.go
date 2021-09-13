// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/docker/registry (interfaces: Registry,RegistryInternal,Matcher,Initializer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	docker "github.com/juju/juju/docker"
	tools "github.com/juju/juju/tools"
)

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRegistry) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRegistryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegistry)(nil).Close))
}

// ImageRepoDetails mocks base method.
func (m *MockRegistry) ImageRepoDetails() docker.ImageRepoDetails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageRepoDetails")
	ret0, _ := ret[0].(docker.ImageRepoDetails)
	return ret0
}

// ImageRepoDetails indicates an expected call of ImageRepoDetails.
func (mr *MockRegistryMockRecorder) ImageRepoDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageRepoDetails", reflect.TypeOf((*MockRegistry)(nil).ImageRepoDetails))
}

// Ping mocks base method.
func (m *MockRegistry) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockRegistryMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRegistry)(nil).Ping))
}

// Tags mocks base method.
func (m *MockRegistry) Tags(arg0 string) (tools.Versions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags", arg0)
	ret0, _ := ret[0].(tools.Versions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tags indicates an expected call of Tags.
func (mr *MockRegistryMockRecorder) Tags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockRegistry)(nil).Tags), arg0)
}

// MockRegistryInternal is a mock of RegistryInternal interface.
type MockRegistryInternal struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryInternalMockRecorder
}

// MockRegistryInternalMockRecorder is the mock recorder for MockRegistryInternal.
type MockRegistryInternalMockRecorder struct {
	mock *MockRegistryInternal
}

// NewMockRegistryInternal creates a new mock instance.
func NewMockRegistryInternal(ctrl *gomock.Controller) *MockRegistryInternal {
	mock := &MockRegistryInternal{ctrl: ctrl}
	mock.recorder = &MockRegistryInternalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistryInternal) EXPECT() *MockRegistryInternalMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRegistryInternal) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRegistryInternalMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegistryInternal)(nil).Close))
}

// DecideBaseURL mocks base method.
func (m *MockRegistryInternal) DecideBaseURL() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecideBaseURL")
	ret0, _ := ret[0].(error)
	return ret0
}

// DecideBaseURL indicates an expected call of DecideBaseURL.
func (mr *MockRegistryInternalMockRecorder) DecideBaseURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecideBaseURL", reflect.TypeOf((*MockRegistryInternal)(nil).DecideBaseURL))
}

// ImageRepoDetails mocks base method.
func (m *MockRegistryInternal) ImageRepoDetails() docker.ImageRepoDetails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageRepoDetails")
	ret0, _ := ret[0].(docker.ImageRepoDetails)
	return ret0
}

// ImageRepoDetails indicates an expected call of ImageRepoDetails.
func (mr *MockRegistryInternalMockRecorder) ImageRepoDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageRepoDetails", reflect.TypeOf((*MockRegistryInternal)(nil).ImageRepoDetails))
}

// Match mocks base method.
func (m *MockRegistryInternal) Match() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match.
func (mr *MockRegistryInternalMockRecorder) Match() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockRegistryInternal)(nil).Match))
}

// Ping mocks base method.
func (m *MockRegistryInternal) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockRegistryInternalMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRegistryInternal)(nil).Ping))
}

// Tags mocks base method.
func (m *MockRegistryInternal) Tags(arg0 string) (tools.Versions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags", arg0)
	ret0, _ := ret[0].(tools.Versions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tags indicates an expected call of Tags.
func (mr *MockRegistryInternalMockRecorder) Tags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockRegistryInternal)(nil).Tags), arg0)
}

// WrapTransport mocks base method.
func (m *MockRegistryInternal) WrapTransport() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WrapTransport")
	ret0, _ := ret[0].(error)
	return ret0
}

// WrapTransport indicates an expected call of WrapTransport.
func (mr *MockRegistryInternalMockRecorder) WrapTransport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WrapTransport", reflect.TypeOf((*MockRegistryInternal)(nil).WrapTransport))
}

// MockMatcher is a mock of Matcher interface.
type MockMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMockRecorder
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher.
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance.
func NewMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// Match mocks base method.
func (m *MockMatcher) Match() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match.
func (mr *MockMatcherMockRecorder) Match() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockMatcher)(nil).Match))
}

// MockInitializer is a mock of Initializer interface.
type MockInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockInitializerMockRecorder
}

// MockInitializerMockRecorder is the mock recorder for MockInitializer.
type MockInitializerMockRecorder struct {
	mock *MockInitializer
}

// NewMockInitializer creates a new mock instance.
func NewMockInitializer(ctrl *gomock.Controller) *MockInitializer {
	mock := &MockInitializer{ctrl: ctrl}
	mock.recorder = &MockInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInitializer) EXPECT() *MockInitializerMockRecorder {
	return m.recorder
}

// DecideBaseURL mocks base method.
func (m *MockInitializer) DecideBaseURL() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecideBaseURL")
	ret0, _ := ret[0].(error)
	return ret0
}

// DecideBaseURL indicates an expected call of DecideBaseURL.
func (mr *MockInitializerMockRecorder) DecideBaseURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecideBaseURL", reflect.TypeOf((*MockInitializer)(nil).DecideBaseURL))
}

// Ping mocks base method.
func (m *MockInitializer) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockInitializerMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockInitializer)(nil).Ping))
}

// WrapTransport mocks base method.
func (m *MockInitializer) WrapTransport() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WrapTransport")
	ret0, _ := ret[0].(error)
	return ret0
}

// WrapTransport indicates an expected call of WrapTransport.
func (mr *MockInitializerMockRecorder) WrapTransport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WrapTransport", reflect.TypeOf((*MockInitializer)(nil).WrapTransport))
}