// Copyright 2015 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package containerinit_test

import (
	"strings"
	stdtesting "testing"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/network"
	"github.com/juju/juju/internal/cloudconfig/cloudinit"
	"github.com/juju/juju/internal/cloudconfig/containerinit"
	"github.com/juju/juju/internal/container"
	containertesting "github.com/juju/juju/internal/container/testing"
	"github.com/juju/juju/internal/testing"
)

func Test(t *stdtesting.T) {
	gc.TestingT(t)
}

type UserDataSuite struct {
	testing.BaseSuite
}

var _ = gc.Suite(&UserDataSuite{})

func (s *UserDataSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)
}

func CloudInitDataExcludingOutputSection(data string) []string {
	// Extract the "#cloud-config" header and all lines between
	// from the "bootcmd" section up to (but not including) the
	// "output" sections to match against expected. But we cannot
	// possibly handle all the /other/ output that may be added by
	// CloudInitUserData() in the future, so we also truncate at
	// the first runcmd which now happens to include the runcmd's
	// added for raising the network interfaces captured in
	// expectedFallbackUserData. However, the other tests above do
	// check for that output.

	var linesToMatch []string
	seenBootcmd := false
	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, "#cloud-config") {
			linesToMatch = append(linesToMatch, line)
			continue
		}

		if strings.HasPrefix(line, "bootcmd:") {
			seenBootcmd = true
		}

		if strings.HasPrefix(line, "output:") && seenBootcmd {
			break
		}

		if seenBootcmd {
			linesToMatch = append(linesToMatch, line)
		}
	}

	return linesToMatch
}

// TestCloudInitUserDataNoNetworkConfig tests that no network-interfaces, or
// related data, appear in user-data when no networkConfig is passed to
// CloudInitUserData.
func (s *UserDataSuite) TestCloudInitUserDataNoNetworkConfig(c *gc.C) {
	instanceConfig, err := containertesting.MockMachineConfig("1/lxd/0")
	c.Assert(err, jc.ErrorIsNil)

	cfg, err := cloudinit.New(instanceConfig.Base.OS)
	c.Assert(err, jc.ErrorIsNil)

	data, err := containerinit.CloudInitUserData(cfg, instanceConfig, nil)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(data, gc.NotNil)

	linesToMatch := CloudInitDataExcludingOutputSection(string(data))

	c.Assert(strings.Join(linesToMatch, "\n"), gc.Equals, "#cloud-config")
}

// TestCloudInitUserDataSomeNetworkConfig tests that the data generated by
// cloudinit.AddNetworkConfig is applied properly.
func (s *UserDataSuite) TestCloudInitUserDataSomeNetworkConfig(c *gc.C) {
	instanceConfig, err := containertesting.MockMachineConfig("1/lxd/0")
	c.Assert(err, jc.ErrorIsNil)

	nics := network.InterfaceInfos{{
		InterfaceName: "eth0",
		InterfaceType: network.EthernetDevice,
		ConfigType:    network.ConfigDHCP,
	}}

	cfg, err := cloudinit.New(instanceConfig.Base.OS)
	c.Assert(err, jc.ErrorIsNil)

	data, err := containerinit.CloudInitUserData(cfg, instanceConfig, container.BridgeNetworkConfig(0, nics))
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(data, gc.NotNil)

	linesToMatch := CloudInitDataExcludingOutputSection(string(data))

	// We just check first two lines, rest is tested deeper.
	c.Assert(linesToMatch[0], gc.Equals, "#cloud-config")
	c.Assert(linesToMatch[1], gc.Equals, "bootcmd:")

}
