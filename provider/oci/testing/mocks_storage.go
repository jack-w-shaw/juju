// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/oci (interfaces: StorageClient)
//
// Generated by this command:
//
//	mockgen -package testing -destination testing/mocks_storage.go -write_package_comment=false github.com/juju/juju/provider/oci StorageClient
package testing

import (
	context "context"
	reflect "reflect"

	core "github.com/oracle/oci-go-sdk/v47/core"
	gomock "go.uber.org/mock/gomock"
)

// MockStorageClient is a mock of StorageClient interface.
type MockStorageClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClientMockRecorder
}

// MockStorageClientMockRecorder is the mock recorder for MockStorageClient.
type MockStorageClientMockRecorder struct {
	mock *MockStorageClient
}

// NewMockStorageClient creates a new mock instance.
func NewMockStorageClient(ctrl *gomock.Controller) *MockStorageClient {
	mock := &MockStorageClient{ctrl: ctrl}
	mock.recorder = &MockStorageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageClient) EXPECT() *MockStorageClientMockRecorder {
	return m.recorder
}

// CreateVolume mocks base method.
func (m *MockStorageClient) CreateVolume(arg0 context.Context, arg1 core.CreateVolumeRequest) (core.CreateVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1)
	ret0, _ := ret[0].(core.CreateVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockStorageClientMockRecorder) CreateVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockStorageClient)(nil).CreateVolume), arg0, arg1)
}

// DeleteVolume mocks base method.
func (m *MockStorageClient) DeleteVolume(arg0 context.Context, arg1 core.DeleteVolumeRequest) (core.DeleteVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(core.DeleteVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockStorageClientMockRecorder) DeleteVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockStorageClient)(nil).DeleteVolume), arg0, arg1)
}

// GetVolume mocks base method.
func (m *MockStorageClient) GetVolume(arg0 context.Context, arg1 core.GetVolumeRequest) (core.GetVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", arg0, arg1)
	ret0, _ := ret[0].(core.GetVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolume indicates an expected call of GetVolume.
func (mr *MockStorageClientMockRecorder) GetVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockStorageClient)(nil).GetVolume), arg0, arg1)
}

// ListVolumes mocks base method.
func (m *MockStorageClient) ListVolumes(arg0 context.Context, arg1 *string) ([]core.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumes", arg0, arg1)
	ret0, _ := ret[0].([]core.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumes indicates an expected call of ListVolumes.
func (mr *MockStorageClientMockRecorder) ListVolumes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockStorageClient)(nil).ListVolumes), arg0, arg1)
}

// UpdateVolume mocks base method.
func (m *MockStorageClient) UpdateVolume(arg0 context.Context, arg1 core.UpdateVolumeRequest) (core.UpdateVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVolume", arg0, arg1)
	ret0, _ := ret[0].(core.UpdateVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVolume indicates an expected call of UpdateVolume.
func (mr *MockStorageClientMockRecorder) UpdateVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVolume", reflect.TypeOf((*MockStorageClient)(nil).UpdateVolume), arg0, arg1)
}
