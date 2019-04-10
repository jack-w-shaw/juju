// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package provider_test

import (
	"fmt"

	"github.com/juju/collections/set"
	"github.com/juju/loggo"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/caas"
	"github.com/juju/juju/caas/kubernetes/provider"
	"github.com/juju/juju/cloud"
	jujucloud "github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
)

var (
	_ = gc.Suite(&cloudSuite{})
)

type cloudSuite struct {
	fakeBroker fakeK8sClusterMetadataChecker
}

var defaultK8sCloud = jujucloud.Cloud{
	Name:           caas.Microk8s,
	Endpoint:       "http://1.1.1.1:8080",
	Type:           cloud.CloudTypeCAAS,
	AuthTypes:      []cloud.AuthType{cloud.UserPassAuthType},
	CACertificates: []string{""},
}

var defaultClusterMetadata = &caas.ClusterMetadata{
	Cloud:                caas.Microk8s,
	Regions:              set.NewStrings(caas.Microk8sRegion),
	OperatorStorageClass: &caas.StorageProvisioner{Name: "operator-sc"},
}

func getDefaultCredential() cloud.Credential {
	defaultCredential := cloud.NewCredential(cloud.UserPassAuthType, map[string]string{"username": "admin", "password": ""})
	defaultCredential.Label = "kubernetes credential \"admin\""
	return defaultCredential
}

func (s *cloudSuite) SetUpTest(c *gc.C) {
	var logger loggo.Logger
	s.fakeBroker = fakeK8sClusterMetadataChecker{CallMocker: testing.NewCallMocker(logger)}
}

func (s *cloudSuite) TestFinalizeCloudNotMicrok8s(c *gc.C) {
	notK8sCloud := jujucloud.Cloud{}
	p := provider.NewProviderWithFakes(
		dummyRunner{},
		getterFunc(builtinCloudRet{}),
		func(environs.OpenParams) (caas.ClusterMetadataChecker, error) { return &s.fakeBroker, nil })
	cloudFinalizer := p.(environs.CloudFinalizer)

	var ctx mockContext
	cloud, err := cloudFinalizer.FinalizeCloud(&ctx, notK8sCloud)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(cloud, jc.DeepEquals, notK8sCloud)
}

func (s *cloudSuite) TestFinalizeCloudMicrok8s(c *gc.C) {
	p := s.getProvider()
	cloudFinalizer := p.(environs.CloudFinalizer)

	var ctx mockContext
	cloud, err := cloudFinalizer.FinalizeCloud(&ctx, defaultK8sCloud)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(cloud, jc.DeepEquals, jujucloud.Cloud{
		Name:            caas.Microk8s,
		Type:            jujucloud.CloudTypeCAAS,
		AuthTypes:       []jujucloud.AuthType{jujucloud.UserPassAuthType},
		CACertificates:  []string{""},
		Endpoint:        "http://1.1.1.1:8080",
		HostCloudRegion: fmt.Sprintf("%s/%s", caas.Microk8s, caas.Microk8sRegion),
		Config:          map[string]interface{}{"operator-storage": "operator-sc", "workload-storage": ""},
		Regions:         []jujucloud.Region{{Name: caas.Microk8sRegion, Endpoint: "http://1.1.1.1:8080"}},
	})
}

func (s *cloudSuite) getProvider() caas.ContainerEnvironProvider {
	s.fakeBroker.Call("GetClusterMetadata").Returns(defaultClusterMetadata, nil)
	s.fakeBroker.Call("CheckDefaultWorkloadStorage").Returns(nil)
	return provider.NewProviderWithFakes(
		dummyRunner{},
		getterFunc(builtinCloudRet{cloud: defaultK8sCloud, credential: getDefaultCredential(), err: nil}),
		func(environs.OpenParams) (caas.ClusterMetadataChecker, error) { return &s.fakeBroker, nil },
	)
}

type mockContext struct {
	testing.Stub
}

func (c *mockContext) Verbosef(f string, args ...interface{}) {
	c.MethodCall(c, "Verbosef", f, args)
}

type fakeK8sClusterMetadataChecker struct {
	*testing.CallMocker
	caas.ClusterMetadataChecker
}

func (api *fakeK8sClusterMetadataChecker) GetClusterMetadata(storageClass string) (result *caas.ClusterMetadata, err error) {
	results := api.MethodCall(api, "GetClusterMetadata")
	return results[0].(*caas.ClusterMetadata), testing.TypeAssertError(results[1])
}

func (api *fakeK8sClusterMetadataChecker) CheckDefaultWorkloadStorage(cluster string, storageProvisioner *caas.StorageProvisioner) error {
	results := api.MethodCall(api, "CheckDefaultWorkloadStorage")
	return testing.TypeAssertError(results[0])
}

func (api *fakeK8sClusterMetadataChecker) EnsureStorageProvisioner(cfg caas.StorageProvisioner) (*caas.StorageProvisioner, error) {
	results := api.MethodCall(api, "EnsureStorageProvisioner")
	return results[0].(*caas.StorageProvisioner), testing.TypeAssertError(results[1])
}
