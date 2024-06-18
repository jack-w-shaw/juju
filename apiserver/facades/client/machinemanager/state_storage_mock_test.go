// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state/binarystorage (interfaces: StorageCloser)
//
// Generated by this command:
//
//	mockgen -typed -package machinemanager_test -destination state_storage_mock_test.go github.com/juju/juju/state/binarystorage StorageCloser
//

// Package machinemanager_test is a generated GoMock package.
package machinemanager_test

import (
	context "context"
	io "io"
	reflect "reflect"

	binarystorage "github.com/juju/juju/state/binarystorage"
	gomock "go.uber.org/mock/gomock"
)

// MockStorageCloser is a mock of StorageCloser interface.
type MockStorageCloser struct {
	ctrl     *gomock.Controller
	recorder *MockStorageCloserMockRecorder
}

// MockStorageCloserMockRecorder is the mock recorder for MockStorageCloser.
type MockStorageCloserMockRecorder struct {
	mock *MockStorageCloser
}

// NewMockStorageCloser creates a new mock instance.
func NewMockStorageCloser(ctrl *gomock.Controller) *MockStorageCloser {
	mock := &MockStorageCloser{ctrl: ctrl}
	mock.recorder = &MockStorageCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageCloser) EXPECT() *MockStorageCloserMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStorageCloser) Add(arg0 context.Context, arg1 io.Reader, arg2 binarystorage.Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStorageCloserMockRecorder) Add(arg0, arg1, arg2 any) *MockStorageCloserAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStorageCloser)(nil).Add), arg0, arg1, arg2)
	return &MockStorageCloserAddCall{Call: call}
}

// MockStorageCloserAddCall wrap *gomock.Call
type MockStorageCloserAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageCloserAddCall) Return(arg0 error) *MockStorageCloserAddCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageCloserAddCall) Do(f func(context.Context, io.Reader, binarystorage.Metadata) error) *MockStorageCloserAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageCloserAddCall) DoAndReturn(f func(context.Context, io.Reader, binarystorage.Metadata) error) *MockStorageCloserAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllMetadata mocks base method.
func (m *MockStorageCloser) AllMetadata() ([]binarystorage.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllMetadata")
	ret0, _ := ret[0].([]binarystorage.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllMetadata indicates an expected call of AllMetadata.
func (mr *MockStorageCloserMockRecorder) AllMetadata() *MockStorageCloserAllMetadataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllMetadata", reflect.TypeOf((*MockStorageCloser)(nil).AllMetadata))
	return &MockStorageCloserAllMetadataCall{Call: call}
}

// MockStorageCloserAllMetadataCall wrap *gomock.Call
type MockStorageCloserAllMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageCloserAllMetadataCall) Return(arg0 []binarystorage.Metadata, arg1 error) *MockStorageCloserAllMetadataCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageCloserAllMetadataCall) Do(f func() ([]binarystorage.Metadata, error)) *MockStorageCloserAllMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageCloserAllMetadataCall) DoAndReturn(f func() ([]binarystorage.Metadata, error)) *MockStorageCloserAllMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockStorageCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageCloserMockRecorder) Close() *MockStorageCloserCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageCloser)(nil).Close))
	return &MockStorageCloserCloseCall{Call: call}
}

// MockStorageCloserCloseCall wrap *gomock.Call
type MockStorageCloserCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageCloserCloseCall) Return(arg0 error) *MockStorageCloserCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageCloserCloseCall) Do(f func() error) *MockStorageCloserCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageCloserCloseCall) DoAndReturn(f func() error) *MockStorageCloserCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Metadata mocks base method.
func (m *MockStorageCloser) Metadata(arg0 string) (binarystorage.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata", arg0)
	ret0, _ := ret[0].(binarystorage.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata.
func (mr *MockStorageCloserMockRecorder) Metadata(arg0 any) *MockStorageCloserMetadataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockStorageCloser)(nil).Metadata), arg0)
	return &MockStorageCloserMetadataCall{Call: call}
}

// MockStorageCloserMetadataCall wrap *gomock.Call
type MockStorageCloserMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageCloserMetadataCall) Return(arg0 binarystorage.Metadata, arg1 error) *MockStorageCloserMetadataCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageCloserMetadataCall) Do(f func(string) (binarystorage.Metadata, error)) *MockStorageCloserMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageCloserMetadataCall) DoAndReturn(f func(string) (binarystorage.Metadata, error)) *MockStorageCloserMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Open mocks base method.
func (m *MockStorageCloser) Open(arg0 context.Context, arg1 string) (binarystorage.Metadata, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(binarystorage.Metadata)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Open indicates an expected call of Open.
func (mr *MockStorageCloserMockRecorder) Open(arg0, arg1 any) *MockStorageCloserOpenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockStorageCloser)(nil).Open), arg0, arg1)
	return &MockStorageCloserOpenCall{Call: call}
}

// MockStorageCloserOpenCall wrap *gomock.Call
type MockStorageCloserOpenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageCloserOpenCall) Return(arg0 binarystorage.Metadata, arg1 io.ReadCloser, arg2 error) *MockStorageCloserOpenCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageCloserOpenCall) Do(f func(context.Context, string) (binarystorage.Metadata, io.ReadCloser, error)) *MockStorageCloserOpenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageCloserOpenCall) DoAndReturn(f func(context.Context, string) (binarystorage.Metadata, io.ReadCloser, error)) *MockStorageCloserOpenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}