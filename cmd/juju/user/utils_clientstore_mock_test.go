// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/jujuclient (interfaces: ClientStore)
//
// Generated by this command:
//
//	mockgen -package user_test -destination utils_clientstore_mock_test.go github.com/juju/juju/jujuclient ClientStore
//

// Package user_test is a generated GoMock package.
package user_test

import (
	reflect "reflect"

	cloud "github.com/juju/juju/cloud"
	jujuclient "github.com/juju/juju/jujuclient"
	gomock "go.uber.org/mock/gomock"
)

// MockClientStore is a mock of ClientStore interface.
type MockClientStore struct {
	ctrl     *gomock.Controller
	recorder *MockClientStoreMockRecorder
}

// MockClientStoreMockRecorder is the mock recorder for MockClientStore.
type MockClientStoreMockRecorder struct {
	mock *MockClientStore
}

// NewMockClientStore creates a new mock instance.
func NewMockClientStore(ctrl *gomock.Controller) *MockClientStore {
	mock := &MockClientStore{ctrl: ctrl}
	mock.recorder = &MockClientStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientStore) EXPECT() *MockClientStoreMockRecorder {
	return m.recorder
}

// AccountDetails mocks base method.
func (m *MockClientStore) AccountDetails(arg0 string) (*jujuclient.AccountDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountDetails", arg0)
	ret0, _ := ret[0].(*jujuclient.AccountDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountDetails indicates an expected call of AccountDetails.
func (mr *MockClientStoreMockRecorder) AccountDetails(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountDetails", reflect.TypeOf((*MockClientStore)(nil).AccountDetails), arg0)
}

// AddController mocks base method.
func (m *MockClientStore) AddController(arg0 string, arg1 jujuclient.ControllerDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddController indicates an expected call of AddController.
func (mr *MockClientStoreMockRecorder) AddController(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddController", reflect.TypeOf((*MockClientStore)(nil).AddController), arg0, arg1)
}

// AllControllers mocks base method.
func (m *MockClientStore) AllControllers() (map[string]jujuclient.ControllerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllControllers")
	ret0, _ := ret[0].(map[string]jujuclient.ControllerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllControllers indicates an expected call of AllControllers.
func (mr *MockClientStoreMockRecorder) AllControllers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllControllers", reflect.TypeOf((*MockClientStore)(nil).AllControllers))
}

// AllCredentials mocks base method.
func (m *MockClientStore) AllCredentials() (map[string]cloud.CloudCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllCredentials")
	ret0, _ := ret[0].(map[string]cloud.CloudCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllCredentials indicates an expected call of AllCredentials.
func (mr *MockClientStoreMockRecorder) AllCredentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllCredentials", reflect.TypeOf((*MockClientStore)(nil).AllCredentials))
}

// AllModels mocks base method.
func (m *MockClientStore) AllModels(arg0 string) (map[string]jujuclient.ModelDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModels", arg0)
	ret0, _ := ret[0].(map[string]jujuclient.ModelDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModels indicates an expected call of AllModels.
func (mr *MockClientStoreMockRecorder) AllModels(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModels", reflect.TypeOf((*MockClientStore)(nil).AllModels), arg0)
}

// BootstrapConfigForController mocks base method.
func (m *MockClientStore) BootstrapConfigForController(arg0 string) (*jujuclient.BootstrapConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapConfigForController", arg0)
	ret0, _ := ret[0].(*jujuclient.BootstrapConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BootstrapConfigForController indicates an expected call of BootstrapConfigForController.
func (mr *MockClientStoreMockRecorder) BootstrapConfigForController(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapConfigForController", reflect.TypeOf((*MockClientStore)(nil).BootstrapConfigForController), arg0)
}

// ControllerByAPIEndpoints mocks base method.
func (m *MockClientStore) ControllerByAPIEndpoints(arg0 ...string) (*jujuclient.ControllerDetails, string, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ControllerByAPIEndpoints", varargs...)
	ret0, _ := ret[0].(*jujuclient.ControllerDetails)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ControllerByAPIEndpoints indicates an expected call of ControllerByAPIEndpoints.
func (mr *MockClientStoreMockRecorder) ControllerByAPIEndpoints(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerByAPIEndpoints", reflect.TypeOf((*MockClientStore)(nil).ControllerByAPIEndpoints), arg0...)
}

// ControllerByName mocks base method.
func (m *MockClientStore) ControllerByName(arg0 string) (*jujuclient.ControllerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerByName", arg0)
	ret0, _ := ret[0].(*jujuclient.ControllerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerByName indicates an expected call of ControllerByName.
func (mr *MockClientStoreMockRecorder) ControllerByName(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerByName", reflect.TypeOf((*MockClientStore)(nil).ControllerByName), arg0)
}

// CookieJar mocks base method.
func (m *MockClientStore) CookieJar(arg0 string) (jujuclient.CookieJar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CookieJar", arg0)
	ret0, _ := ret[0].(jujuclient.CookieJar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CookieJar indicates an expected call of CookieJar.
func (mr *MockClientStoreMockRecorder) CookieJar(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CookieJar", reflect.TypeOf((*MockClientStore)(nil).CookieJar), arg0)
}

// CredentialForCloud mocks base method.
func (m *MockClientStore) CredentialForCloud(arg0 string) (*cloud.CloudCredential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialForCloud", arg0)
	ret0, _ := ret[0].(*cloud.CloudCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredentialForCloud indicates an expected call of CredentialForCloud.
func (mr *MockClientStoreMockRecorder) CredentialForCloud(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialForCloud", reflect.TypeOf((*MockClientStore)(nil).CredentialForCloud), arg0)
}

// CurrentController mocks base method.
func (m *MockClientStore) CurrentController() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentController")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentController indicates an expected call of CurrentController.
func (mr *MockClientStoreMockRecorder) CurrentController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentController", reflect.TypeOf((*MockClientStore)(nil).CurrentController))
}

// CurrentModel mocks base method.
func (m *MockClientStore) CurrentModel(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentModel", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentModel indicates an expected call of CurrentModel.
func (mr *MockClientStoreMockRecorder) CurrentModel(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentModel", reflect.TypeOf((*MockClientStore)(nil).CurrentModel), arg0)
}

// ModelByName mocks base method.
func (m *MockClientStore) ModelByName(arg0, arg1 string) (*jujuclient.ModelDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelByName", arg0, arg1)
	ret0, _ := ret[0].(*jujuclient.ModelDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelByName indicates an expected call of ModelByName.
func (mr *MockClientStoreMockRecorder) ModelByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelByName", reflect.TypeOf((*MockClientStore)(nil).ModelByName), arg0, arg1)
}

// RemoveAccount mocks base method.
func (m *MockClientStore) RemoveAccount(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAccount indicates an expected call of RemoveAccount.
func (mr *MockClientStoreMockRecorder) RemoveAccount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccount", reflect.TypeOf((*MockClientStore)(nil).RemoveAccount), arg0)
}

// RemoveController mocks base method.
func (m *MockClientStore) RemoveController(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveController", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveController indicates an expected call of RemoveController.
func (mr *MockClientStoreMockRecorder) RemoveController(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveController", reflect.TypeOf((*MockClientStore)(nil).RemoveController), arg0)
}

// RemoveModel mocks base method.
func (m *MockClientStore) RemoveModel(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveModel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveModel indicates an expected call of RemoveModel.
func (mr *MockClientStoreMockRecorder) RemoveModel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveModel", reflect.TypeOf((*MockClientStore)(nil).RemoveModel), arg0, arg1)
}

// SetCurrentController mocks base method.
func (m *MockClientStore) SetCurrentController(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurrentController", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCurrentController indicates an expected call of SetCurrentController.
func (mr *MockClientStoreMockRecorder) SetCurrentController(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentController", reflect.TypeOf((*MockClientStore)(nil).SetCurrentController), arg0)
}

// SetCurrentModel mocks base method.
func (m *MockClientStore) SetCurrentModel(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurrentModel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCurrentModel indicates an expected call of SetCurrentModel.
func (mr *MockClientStoreMockRecorder) SetCurrentModel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentModel", reflect.TypeOf((*MockClientStore)(nil).SetCurrentModel), arg0, arg1)
}

// SetModels mocks base method.
func (m *MockClientStore) SetModels(arg0 string, arg1 map[string]jujuclient.ModelDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModels", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModels indicates an expected call of SetModels.
func (mr *MockClientStoreMockRecorder) SetModels(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModels", reflect.TypeOf((*MockClientStore)(nil).SetModels), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockClientStore) UpdateAccount(arg0 string, arg1 jujuclient.AccountDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockClientStoreMockRecorder) UpdateAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockClientStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateBootstrapConfig mocks base method.
func (m *MockClientStore) UpdateBootstrapConfig(arg0 string, arg1 jujuclient.BootstrapConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBootstrapConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBootstrapConfig indicates an expected call of UpdateBootstrapConfig.
func (mr *MockClientStoreMockRecorder) UpdateBootstrapConfig(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBootstrapConfig", reflect.TypeOf((*MockClientStore)(nil).UpdateBootstrapConfig), arg0, arg1)
}

// UpdateController mocks base method.
func (m *MockClientStore) UpdateController(arg0 string, arg1 jujuclient.ControllerDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateController indicates an expected call of UpdateController.
func (mr *MockClientStoreMockRecorder) UpdateController(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateController", reflect.TypeOf((*MockClientStore)(nil).UpdateController), arg0, arg1)
}

// UpdateCredential mocks base method.
func (m *MockClientStore) UpdateCredential(arg0 string, arg1 cloud.CloudCredential) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCredential", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCredential indicates an expected call of UpdateCredential.
func (mr *MockClientStoreMockRecorder) UpdateCredential(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCredential", reflect.TypeOf((*MockClientStore)(nil).UpdateCredential), arg0, arg1)
}

// UpdateModel mocks base method.
func (m *MockClientStore) UpdateModel(arg0, arg1 string, arg2 jujuclient.ModelDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateModel", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateModel indicates an expected call of UpdateModel.
func (mr *MockClientStoreMockRecorder) UpdateModel(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateModel", reflect.TypeOf((*MockClientStore)(nil).UpdateModel), arg0, arg1, arg2)
}
