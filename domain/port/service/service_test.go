// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"
	"sort"

	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/application"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/unit"
	domain "github.com/juju/juju/domain"
	"github.com/juju/juju/domain/port"
	domaintesting "github.com/juju/juju/domain/testing"
	"github.com/juju/juju/internal/uuid"
)

type serviceSuite struct {
	st *MockState

	wildcardEpUUID port.UUID
	epUUID1        port.UUID
	epUUID2        port.UUID
	epUUID3        port.UUID

	portRangeUUID1 port.UUID
	portRangeUUID2 port.UUID
	portRangeUUID3 port.UUID

	unitUUID1 unit.UUID
	unitUUID2 unit.UUID

	machineUUID string

	appUUID application.ID
}

var _ = gc.Suite(&serviceSuite{})

func (s *serviceSuite) SetUpTest(c *gc.C) {
	var err error

	s.wildcardEpUUID, err = port.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	s.epUUID1, err = port.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	s.epUUID2, err = port.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	s.epUUID3, err = port.NewUUID()
	c.Assert(err, jc.ErrorIsNil)

	s.portRangeUUID1, err = port.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	s.portRangeUUID2, err = port.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	s.portRangeUUID3, err = port.NewUUID()
	c.Assert(err, jc.ErrorIsNil)

	s.unitUUID1, err = unit.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	s.unitUUID2, err = unit.NewUUID()
	c.Assert(err, jc.ErrorIsNil)

	s.machineUUID = uuid.MustNewUUID().String()

	s.appUUID, err = application.NewID()
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.st = NewMockState(ctrl)
	s.st.EXPECT().RunAtomic(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx domain.AtomicContext) error) error {
		return fn(domaintesting.NewAtomicContext(ctx))
	}).AnyTimes()

	return ctrl
}

func (s *serviceSuite) TestGetUnitOpenedPorts(c *gc.C) {
	defer s.setupMocks(c).Finish()

	grp := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
		"ep2": {
			network.MustParsePortRange("8000-9000/udp"),
		},
	}

	s.st.EXPECT().GetUnitOpenedPorts(gomock.Any(), s.unitUUID1).Return(grp, nil)

	srv := NewService(s.st)
	res, err := srv.GetUnitOpenedPorts(context.Background(), s.unitUUID1)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(res, gc.DeepEquals, grp)
}

func (s *serviceSuite) TestGetMachineOpenedPorts(c *gc.C) {
	defer s.setupMocks(c).Finish()

	grp := map[unit.UUID]network.GroupedPortRanges{
		s.unitUUID1: {
			"ep1": {
				network.MustParsePortRange("80/tcp"),
				network.MustParsePortRange("443/tcp"),
			},
			"ep2": {
				network.MustParsePortRange("8000-9000/udp"),
			},
		},
		s.unitUUID2: {
			"ep3": {
				network.MustParsePortRange("8080/tcp"),
			},
		},
	}

	s.st.EXPECT().GetMachineOpenedPorts(gomock.Any(), s.machineUUID).Return(grp, nil)

	srv := NewService(s.st)
	res, err := srv.GetMachineOpenedPorts(context.Background(), s.machineUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(res, gc.DeepEquals, grp)
}

func (s *serviceSuite) TestGetApplicationOpenedPorts(c *gc.C) {
	defer s.setupMocks(c).Finish()

	openedPorts := port.UnitEndpointPortRanges{
		{Endpoint: "ep1", UnitUUID: s.unitUUID1, PortRange: network.MustParsePortRange("80/tcp")},
		{Endpoint: "ep1", UnitUUID: s.unitUUID1, PortRange: network.MustParsePortRange("443/tcp")},
		{Endpoint: "ep2", UnitUUID: s.unitUUID1, PortRange: network.MustParsePortRange("8000-9000/udp")},
		{Endpoint: "ep3", UnitUUID: s.unitUUID2, PortRange: network.MustParsePortRange("8080/tcp")},
	}

	expected := map[unit.UUID]network.GroupedPortRanges{
		s.unitUUID1: {
			"ep1": {
				network.MustParsePortRange("80/tcp"),
				network.MustParsePortRange("443/tcp"),
			},
			"ep2": {
				network.MustParsePortRange("8000-9000/udp"),
			},
		},
		s.unitUUID2: {
			"ep3": {
				network.MustParsePortRange("8080/tcp"),
			},
		},
	}

	s.st.EXPECT().GetApplicationOpenedPorts(gomock.Any(), s.appUUID).Return(openedPorts, nil)

	srv := NewService(s.st)
	res, err := srv.GetApplicationOpenedPorts(context.Background(), s.appUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(res, gc.DeepEquals, expected)
}

func (s *serviceSuite) TestGetApplicationOpenedPortsByEndpoint(c *gc.C) {
	defer s.setupMocks(c).Finish()

	openedPorts := port.UnitEndpointPortRanges{
		{Endpoint: "ep1", UnitUUID: s.unitUUID1, PortRange: network.MustParsePortRange("80/tcp")},
		{Endpoint: "ep1", UnitUUID: s.unitUUID1, PortRange: network.MustParsePortRange("443/tcp")},
		{Endpoint: "ep1", UnitUUID: s.unitUUID2, PortRange: network.MustParsePortRange("8080/tcp")},
		{Endpoint: "ep2", UnitUUID: s.unitUUID1, PortRange: network.MustParsePortRange("8000-8005/udp")},
	}

	s.st.EXPECT().GetApplicationOpenedPorts(gomock.Any(), s.appUUID).Return(openedPorts, nil)

	expected := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
			network.MustParsePortRange("8080/tcp"),
		},
		"ep2": {
			network.MustParsePortRange("8000/udp"),
			network.MustParsePortRange("8001/udp"),
			network.MustParsePortRange("8002/udp"),
			network.MustParsePortRange("8003/udp"),
			network.MustParsePortRange("8004/udp"),
			network.MustParsePortRange("8005/udp"),
		},
	}

	srv := NewService(s.st)
	res, err := srv.GetApplicationOpenedPortsByEndpoint(context.Background(), s.appUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(res, gc.DeepEquals, expected)
}

func (s serviceSuite) TestGetApplicationOpenedPortsByEndpointOverlap(c *gc.C) {
	defer s.setupMocks(c).Finish()

	openedPorts := port.UnitEndpointPortRanges{
		{Endpoint: "ep1", UnitUUID: s.unitUUID1, PortRange: network.MustParsePortRange("80-85/tcp")},
		{Endpoint: "ep1", UnitUUID: s.unitUUID2, PortRange: network.MustParsePortRange("83-88/tcp")},
	}

	s.st.EXPECT().GetApplicationOpenedPorts(gomock.Any(), s.appUUID).Return(openedPorts, nil)

	expected := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("81/tcp"),
			network.MustParsePortRange("82/tcp"),
			network.MustParsePortRange("83/tcp"),
			network.MustParsePortRange("84/tcp"),
			network.MustParsePortRange("85/tcp"),
			network.MustParsePortRange("86/tcp"),
			network.MustParsePortRange("87/tcp"),
			network.MustParsePortRange("88/tcp"),
		},
	}

	srv := NewService(s.st)
	res, err := srv.GetApplicationOpenedPortsByEndpoint(context.Background(), s.appUUID)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(res, gc.DeepEquals, expected)
}

func (s *serviceSuite) TestUpdateUnitPorts(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		"ep1": {
			{UUID: s.portRangeUUID1, PortRange: network.MustParsePortRange("22/tcp")},
			{UUID: s.portRangeUUID2, PortRange: network.MustParsePortRange("23/tcp")},
		},
	}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
		"ep2": {
			network.MustParsePortRange("8000-9000/udp"),
		},
	}

	closePorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("22/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.epUUID1: {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
		s.epUUID2: {
			network.MustParsePortRange("8000-9000/udp"),
		},
	}).Return(nil)
	s.st.EXPECT().RemoveOpenedPorts(domaintesting.IsAtomicContextChecker, []port.UUID{s.portRangeUUID1}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsNoChanges(c *gc.C) {
	srv := NewService(nil)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, network.GroupedPortRanges{"ep1": {}}, network.GroupedPortRanges{})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsClosePortsNotOpen(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{}

	openPorts := network.GroupedPortRanges{}
	closePorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("22/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsOpenedPortsAlreadyOpen(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		"ep1": {
			{UUID: s.portRangeUUID1, PortRange: network.MustParsePortRange("80/tcp")},
		},
	}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
	}
	closePorts := network.GroupedPortRanges{}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.epUUID1: {
			network.MustParsePortRange("443/tcp"),
		},
	}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsSameRangeAcrossEndpoints(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
		{UUID: s.epUUID3, Endpoint: "ep3"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
		"ep2": {
			network.MustParsePortRange("80/tcp"),
		},
		"ep3": {
			network.MustParsePortRange("80/tcp"),
		},
	}
	closePorts := network.GroupedPortRanges{}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.epUUID1: {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
		s.epUUID2: {
			network.MustParsePortRange("80/tcp"),
		},
		s.epUUID3: {
			network.MustParsePortRange("80/tcp"),
		},
	}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsConflict(c *gc.C) {
	srv := NewService(nil)

	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
		"ep2": {
			network.MustParsePortRange("150-250/tcp"),
		},
	}, network.GroupedPortRanges{})
	c.Assert(err, jc.ErrorIs, port.ErrPortRangeConflict)

	err = srv.UpdateUnitPorts(context.Background(), s.unitUUID1, network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
	}, network.GroupedPortRanges{
		"ep2": {
			network.MustParsePortRange("150-250/tcp"),
		},
	})
	c.Assert(err, jc.ErrorIs, port.ErrPortRangeConflict)

	err = srv.UpdateUnitPorts(context.Background(), s.unitUUID1, network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
			network.MustParsePortRange("200/tcp"),
		},
	}, network.GroupedPortRanges{})
	c.Assert(err, jc.ErrorIs, port.ErrPortRangeConflict)
}

func (s *serviceSuite) TestUpdateUnitPortsOpenPortConflictColocated(c *gc.C) {
	defer s.setupMocks(c).Finish()

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{
		network.MustParsePortRange("150-250/tcp"),
	}, nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, network.GroupedPortRanges{})

	c.Assert(err, jc.ErrorIs, port.ErrPortRangeConflict)
}

func (s *serviceSuite) TestUpdateUnitPortsClosePortConflictColocated(c *gc.C) {
	defer s.setupMocks(c).Finish()

	closePorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{
		network.MustParsePortRange("150-250/tcp"),
	}, nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, network.GroupedPortRanges{}, closePorts)

	c.Assert(err, jc.ErrorIs, port.ErrPortRangeConflict)
}

func (s *serviceSuite) TestUpdateUnitPortsOpenWildcard(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
		{UUID: s.epUUID3, Endpoint: "ep3"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		"ep1": {
			{UUID: s.portRangeUUID1, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
		"ep2": {
			{UUID: s.portRangeUUID2, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
	}

	openPorts := network.GroupedPortRanges{
		WildcardEndpoint: {
			network.MustParsePortRange("100-200/tcp"),
		},
	}
	closePorts := network.GroupedPortRanges{}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.wildcardEpUUID: {
			network.MustParsePortRange("100-200/tcp"),
		},
	}).Return(nil)

	closePortsUUIDs := []port.UUID{s.portRangeUUID1, s.portRangeUUID2}
	sort.Slice(closePortsUUIDs, func(i, j int) bool {
		return closePortsUUIDs[i] < closePortsUUIDs[j]
	})
	s.st.EXPECT().RemoveOpenedPorts(domaintesting.IsAtomicContextChecker, closePortsUUIDs).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsOpenPortOnWildcardAndOtherEndpoint(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{}

	openPorts := network.GroupedPortRanges{
		WildcardEndpoint: {
			network.MustParsePortRange("100-200/tcp"),
		},
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
	}

	closePorts := network.GroupedPortRanges{}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.wildcardEpUUID: {
			network.MustParsePortRange("100-200/tcp"),
		},
	}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsOpenPortRangeOpenOnWildcard(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: "wildcard-uuid", Endpoint: WildcardEndpoint},
		{UUID: "ep1-uuid", Endpoint: "ep1"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		WildcardEndpoint: {
			{UUID: "wildcard-uuid", PortRange: network.MustParsePortRange("100-200/tcp")},
		},
	}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
	}
	closePorts := network.GroupedPortRanges{}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsCloseWildcard(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
		{UUID: s.epUUID3, Endpoint: "ep3"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		WildcardEndpoint: {
			{UUID: s.portRangeUUID3, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
		"ep1": {
			{UUID: s.portRangeUUID1, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
		"ep2": {
			{UUID: s.portRangeUUID2, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
	}

	openPorts := network.GroupedPortRanges{}
	closePorts := network.GroupedPortRanges{
		WildcardEndpoint: {
			network.MustParsePortRange("100-200/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)

	closePortsUUIDs := []port.UUID{s.portRangeUUID1, s.portRangeUUID2, s.portRangeUUID3}
	sort.Slice(closePortsUUIDs, func(i, j int) bool {
		return closePortsUUIDs[i] < closePortsUUIDs[j]
	})
	s.st.EXPECT().RemoveOpenedPorts(domaintesting.IsAtomicContextChecker, closePortsUUIDs).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsClosePortRangeOpenOnWildcard(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
		{UUID: s.epUUID3, Endpoint: "ep3"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		WildcardEndpoint: {
			{UUID: s.portRangeUUID1, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
	}

	openPorts := network.GroupedPortRanges{}
	closePorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.epUUID2: {
			network.MustParsePortRange("100-200/tcp"),
		},
		s.epUUID3: {
			network.MustParsePortRange("100-200/tcp"),
		},
	}).Return(nil)
	s.st.EXPECT().RemoveOpenedPorts(domaintesting.IsAtomicContextChecker, []port.UUID{s.portRangeUUID1}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsOpenAndCloseWildcard(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		WildcardEndpoint: {
			{UUID: s.portRangeUUID1, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
		"ep1": {
			{UUID: s.portRangeUUID2, PortRange: network.MustParsePortRange("443/tcp")},
		},
		"ep2": {
			{UUID: s.portRangeUUID3, PortRange: network.MustParsePortRange("444/tcp")},
		},
	}

	openPorts := network.GroupedPortRanges{
		WildcardEndpoint: {
			network.MustParsePortRange("443/tcp"),
		},
	}

	closePorts := network.GroupedPortRanges{
		WildcardEndpoint: {
			network.MustParsePortRange("444/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.wildcardEpUUID: {
			network.MustParsePortRange("443/tcp"),
		},
	}).Return(nil)

	closePortsUUIDs := []port.UUID{s.portRangeUUID2, s.portRangeUUID3}
	sort.Slice(closePortsUUIDs, func(i, j int) bool {
		return closePortsUUIDs[i] < closePortsUUIDs[j]
	})
	s.st.EXPECT().RemoveOpenedPorts(domaintesting.IsAtomicContextChecker, closePortsUUIDs).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortsNeedsToAddSomeEndpoints(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
	}
	newEndpoints := []port.Endpoint{
		{UUID: s.epUUID2, Endpoint: "ep2"},
		{UUID: s.epUUID3, Endpoint: "ep3"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
		"ep2": {
			network.MustParsePortRange("80/tcp"),
		},
		"ep3": {
			network.MustParsePortRange("80/tcp"),
		},
	}
	closePorts := network.GroupedPortRanges{}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return([]network.PortRange{}, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1, []string{"ep2", "ep3"}).Return(newEndpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.epUUID1: {
			network.MustParsePortRange("80/tcp"),
			network.MustParsePortRange("443/tcp"),
		},
		s.epUUID2: {
			network.MustParsePortRange("80/tcp"),
		},
		s.epUUID3: {
			network.MustParsePortRange("80/tcp"),
		},
	}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, closePorts)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortRangeEqualOnColocatedUnit(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
	}

	colocatedPorts := []network.PortRange{
		network.MustParsePortRange("100-150/tcp"),
	}

	currentPorts := map[string][]port.PortRangeWithUUID{}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-150/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(colocatedPorts, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.epUUID1: {
			network.MustParsePortRange("100-150/tcp"),
		},
	}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, network.GroupedPortRanges{})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestUpdateUnitPortRangesConflictsColocated(c *gc.C) {
	defer s.setupMocks(c).Finish()

	colocatedPorts := []network.PortRange{
		network.MustParsePortRange("100-150/tcp"),
	}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(colocatedPorts, nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, network.GroupedPortRanges{})
	c.Assert(err, jc.ErrorIs, port.ErrPortRangeConflict)
}

func (s *serviceSuite) TestUpdateUnitPortRangeEqualOnOtherEndpoint(c *gc.C) {
	defer s.setupMocks(c).Finish()

	endpoints := []port.Endpoint{
		{UUID: s.wildcardEpUUID, Endpoint: WildcardEndpoint},
		{UUID: s.epUUID1, Endpoint: "ep1"},
		{UUID: s.epUUID2, Endpoint: "ep2"},
	}

	currentPorts := map[string][]port.PortRangeWithUUID{
		"ep2": {
			{UUID: s.portRangeUUID1, PortRange: network.MustParsePortRange("100-200/tcp")},
		},
	}

	colocatedPorts := []network.PortRange{network.MustParsePortRange("100-200/tcp")}

	openPorts := network.GroupedPortRanges{
		"ep1": {
			network.MustParsePortRange("100-200/tcp"),
		},
	}

	s.st.EXPECT().GetColocatedOpenedPorts(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(colocatedPorts, nil)
	s.st.EXPECT().GetUnitOpenedPortsWithUUIDs(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(currentPorts, nil)
	s.st.EXPECT().GetEndpoints(domaintesting.IsAtomicContextChecker, s.unitUUID1).Return(endpoints, nil)
	s.st.EXPECT().AddOpenedPorts(domaintesting.IsAtomicContextChecker, map[port.UUID][]network.PortRange{
		s.epUUID1: {
			network.MustParsePortRange("100-200/tcp"),
		},
	}).Return(nil)

	srv := NewService(s.st)
	err := srv.UpdateUnitPorts(context.Background(), s.unitUUID1, openPorts, network.GroupedPortRanges{})
	c.Assert(err, jc.ErrorIsNil)
}
