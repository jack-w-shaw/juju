// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen -typed -package bootstrap -destination charm_mock_test.go github.com/juju/juju/core/charm Repository
//

// Package bootstrap is a generated GoMock package.
package bootstrap

import (
	context "context"
	url "net/url"
	reflect "reflect"

	charm "github.com/juju/juju/core/charm"
	resource "github.com/juju/juju/internal/charm/resource"
	charmhub "github.com/juju/juju/internal/charmhub"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockRepository) Download(arg0 context.Context, arg1 string, arg2 charm.Origin, arg3 string) (charm.Origin, *charmhub.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm.Origin)
	ret1, _ := ret[1].(*charmhub.Digest)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Download indicates an expected call of Download.
func (mr *MockRepositoryMockRecorder) Download(arg0, arg1, arg2, arg3 any) *MockRepositoryDownloadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockRepository)(nil).Download), arg0, arg1, arg2, arg3)
	return &MockRepositoryDownloadCall{Call: call}
}

// MockRepositoryDownloadCall wrap *gomock.Call
type MockRepositoryDownloadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryDownloadCall) Return(arg0 charm.Origin, arg1 *charmhub.Digest, arg2 error) *MockRepositoryDownloadCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryDownloadCall) Do(f func(context.Context, string, charm.Origin, string) (charm.Origin, *charmhub.Digest, error)) *MockRepositoryDownloadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryDownloadCall) DoAndReturn(f func(context.Context, string, charm.Origin, string) (charm.Origin, *charmhub.Digest, error)) *MockRepositoryDownloadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadCharm mocks base method.
func (m *MockRepository) DownloadCharm(arg0 context.Context, arg1 string, arg2 charm.Origin, arg3 string) (charm.CharmArchive, charm.Origin, *charmhub.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadCharm", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm.CharmArchive)
	ret1, _ := ret[1].(charm.Origin)
	ret2, _ := ret[2].(*charmhub.Digest)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// DownloadCharm indicates an expected call of DownloadCharm.
func (mr *MockRepositoryMockRecorder) DownloadCharm(arg0, arg1, arg2, arg3 any) *MockRepositoryDownloadCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadCharm", reflect.TypeOf((*MockRepository)(nil).DownloadCharm), arg0, arg1, arg2, arg3)
	return &MockRepositoryDownloadCharmCall{Call: call}
}

// MockRepositoryDownloadCharmCall wrap *gomock.Call
type MockRepositoryDownloadCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryDownloadCharmCall) Return(arg0 charm.CharmArchive, arg1 charm.Origin, arg2 *charmhub.Digest, arg3 error) *MockRepositoryDownloadCharmCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryDownloadCharmCall) Do(f func(context.Context, string, charm.Origin, string) (charm.CharmArchive, charm.Origin, *charmhub.Digest, error)) *MockRepositoryDownloadCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryDownloadCharmCall) DoAndReturn(f func(context.Context, string, charm.Origin, string) (charm.CharmArchive, charm.Origin, *charmhub.Digest, error)) *MockRepositoryDownloadCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetDownloadURL mocks base method.
func (m *MockRepository) GetDownloadURL(arg0 context.Context, arg1 string, arg2 charm.Origin) (*url.URL, charm.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownloadURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(charm.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDownloadURL indicates an expected call of GetDownloadURL.
func (mr *MockRepositoryMockRecorder) GetDownloadURL(arg0, arg1, arg2 any) *MockRepositoryGetDownloadURLCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadURL", reflect.TypeOf((*MockRepository)(nil).GetDownloadURL), arg0, arg1, arg2)
	return &MockRepositoryGetDownloadURLCall{Call: call}
}

// MockRepositoryGetDownloadURLCall wrap *gomock.Call
type MockRepositoryGetDownloadURLCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryGetDownloadURLCall) Return(arg0 *url.URL, arg1 charm.Origin, arg2 error) *MockRepositoryGetDownloadURLCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryGetDownloadURLCall) Do(f func(context.Context, string, charm.Origin) (*url.URL, charm.Origin, error)) *MockRepositoryGetDownloadURLCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryGetDownloadURLCall) DoAndReturn(f func(context.Context, string, charm.Origin) (*url.URL, charm.Origin, error)) *MockRepositoryGetDownloadURLCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetEssentialMetadata mocks base method.
func (m *MockRepository) GetEssentialMetadata(arg0 context.Context, arg1 ...charm.MetadataRequest) ([]charm.EssentialMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEssentialMetadata", varargs...)
	ret0, _ := ret[0].([]charm.EssentialMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEssentialMetadata indicates an expected call of GetEssentialMetadata.
func (mr *MockRepositoryMockRecorder) GetEssentialMetadata(arg0 any, arg1 ...any) *MockRepositoryGetEssentialMetadataCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEssentialMetadata", reflect.TypeOf((*MockRepository)(nil).GetEssentialMetadata), varargs...)
	return &MockRepositoryGetEssentialMetadataCall{Call: call}
}

// MockRepositoryGetEssentialMetadataCall wrap *gomock.Call
type MockRepositoryGetEssentialMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryGetEssentialMetadataCall) Return(arg0 []charm.EssentialMetadata, arg1 error) *MockRepositoryGetEssentialMetadataCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryGetEssentialMetadataCall) Do(f func(context.Context, ...charm.MetadataRequest) ([]charm.EssentialMetadata, error)) *MockRepositoryGetEssentialMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryGetEssentialMetadataCall) DoAndReturn(f func(context.Context, ...charm.MetadataRequest) ([]charm.EssentialMetadata, error)) *MockRepositoryGetEssentialMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListResources mocks base method.
func (m *MockRepository) ListResources(arg0 context.Context, arg1 string, arg2 charm.Origin) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0, arg1, arg2)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockRepositoryMockRecorder) ListResources(arg0, arg1, arg2 any) *MockRepositoryListResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockRepository)(nil).ListResources), arg0, arg1, arg2)
	return &MockRepositoryListResourcesCall{Call: call}
}

// MockRepositoryListResourcesCall wrap *gomock.Call
type MockRepositoryListResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryListResourcesCall) Return(arg0 []resource.Resource, arg1 error) *MockRepositoryListResourcesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryListResourcesCall) Do(f func(context.Context, string, charm.Origin) ([]resource.Resource, error)) *MockRepositoryListResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryListResourcesCall) DoAndReturn(f func(context.Context, string, charm.Origin) ([]resource.Resource, error)) *MockRepositoryListResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveForDeploy mocks base method.
func (m *MockRepository) ResolveForDeploy(arg0 context.Context, arg1 charm.CharmID) (charm.ResolvedDataForDeploy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveForDeploy", arg0, arg1)
	ret0, _ := ret[0].(charm.ResolvedDataForDeploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveForDeploy indicates an expected call of ResolveForDeploy.
func (mr *MockRepositoryMockRecorder) ResolveForDeploy(arg0, arg1 any) *MockRepositoryResolveForDeployCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveForDeploy", reflect.TypeOf((*MockRepository)(nil).ResolveForDeploy), arg0, arg1)
	return &MockRepositoryResolveForDeployCall{Call: call}
}

// MockRepositoryResolveForDeployCall wrap *gomock.Call
type MockRepositoryResolveForDeployCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryResolveForDeployCall) Return(arg0 charm.ResolvedDataForDeploy, arg1 error) *MockRepositoryResolveForDeployCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryResolveForDeployCall) Do(f func(context.Context, charm.CharmID) (charm.ResolvedDataForDeploy, error)) *MockRepositoryResolveForDeployCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryResolveForDeployCall) DoAndReturn(f func(context.Context, charm.CharmID) (charm.ResolvedDataForDeploy, error)) *MockRepositoryResolveForDeployCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveResources mocks base method.
func (m *MockRepository) ResolveResources(arg0 context.Context, arg1 []resource.Resource, arg2 charm.CharmID) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveResources", arg0, arg1, arg2)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveResources indicates an expected call of ResolveResources.
func (mr *MockRepositoryMockRecorder) ResolveResources(arg0, arg1, arg2 any) *MockRepositoryResolveResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveResources", reflect.TypeOf((*MockRepository)(nil).ResolveResources), arg0, arg1, arg2)
	return &MockRepositoryResolveResourcesCall{Call: call}
}

// MockRepositoryResolveResourcesCall wrap *gomock.Call
type MockRepositoryResolveResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryResolveResourcesCall) Return(arg0 []resource.Resource, arg1 error) *MockRepositoryResolveResourcesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryResolveResourcesCall) Do(f func(context.Context, []resource.Resource, charm.CharmID) ([]resource.Resource, error)) *MockRepositoryResolveResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryResolveResourcesCall) DoAndReturn(f func(context.Context, []resource.Resource, charm.CharmID) ([]resource.Resource, error)) *MockRepositoryResolveResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ResolveWithPreferredChannel mocks base method.
func (m *MockRepository) ResolveWithPreferredChannel(arg0 context.Context, arg1 string, arg2 charm.Origin) (charm.ResolvedData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveWithPreferredChannel", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm.ResolvedData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveWithPreferredChannel indicates an expected call of ResolveWithPreferredChannel.
func (mr *MockRepositoryMockRecorder) ResolveWithPreferredChannel(arg0, arg1, arg2 any) *MockRepositoryResolveWithPreferredChannelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveWithPreferredChannel", reflect.TypeOf((*MockRepository)(nil).ResolveWithPreferredChannel), arg0, arg1, arg2)
	return &MockRepositoryResolveWithPreferredChannelCall{Call: call}
}

// MockRepositoryResolveWithPreferredChannelCall wrap *gomock.Call
type MockRepositoryResolveWithPreferredChannelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRepositoryResolveWithPreferredChannelCall) Return(arg0 charm.ResolvedData, arg1 error) *MockRepositoryResolveWithPreferredChannelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRepositoryResolveWithPreferredChannelCall) Do(f func(context.Context, string, charm.Origin) (charm.ResolvedData, error)) *MockRepositoryResolveWithPreferredChannelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRepositoryResolveWithPreferredChannelCall) DoAndReturn(f func(context.Context, string, charm.Origin) (charm.ResolvedData, error)) *MockRepositoryResolveWithPreferredChannelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
