// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package dbreplaccessor

import (
	"context"
	"testing"
	"time"

	jujutesting "github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4"
	"go.uber.org/goleak"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/logger"
	domaintesting "github.com/juju/juju/domain/schema/testing"
	loggertesting "github.com/juju/juju/internal/logger/testing"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package dbreplaccessor -destination package_mock_test.go github.com/juju/juju/internal/worker/dbreplaccessor DBApp,NodeManager,TrackedDB,Client
//go:generate go run go.uber.org/mock/mockgen -typed -package dbreplaccessor -destination clock_mock_test.go github.com/juju/clock Clock,Timer
//go:generate go run go.uber.org/mock/mockgen -typed -package dbreplaccessor -destination sql_mock_test.go database/sql/driver Driver

func TestPackage(t *testing.T) {
	defer goleak.VerifyNone(t)

	gc.TestingT(t)
}

type baseSuite struct {
	jujutesting.IsolationSuite

	logger logger.Logger

	clock       *MockClock
	timer       *MockTimer
	dbApp       *MockDBApp
	client      *MockClient
	nodeManager *MockNodeManager

	newDBReplWorker func() (TrackedDB, error)
}

func (s *baseSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.clock = NewMockClock(ctrl)
	s.timer = NewMockTimer(ctrl)
	s.dbApp = NewMockDBApp(ctrl)
	s.client = NewMockClient(ctrl)
	s.nodeManager = NewMockNodeManager(ctrl)

	s.logger = loggertesting.WrapCheckLog(c)

	s.newDBReplWorker = nil

	return ctrl
}

func (s *baseSuite) expectClock() {
	s.clock.EXPECT().Now().Return(time.Now()).AnyTimes()
	s.clock.EXPECT().After(gomock.Any()).AnyTimes()
}

func (s *baseSuite) setupTimer(interval time.Duration) chan time.Time {
	s.timer.EXPECT().Stop().MinTimes(1)
	s.clock.EXPECT().NewTimer(interval).Return(s.timer)

	ch := make(chan time.Time)
	s.timer.EXPECT().Chan().Return(ch).AnyTimes()
	return ch
}

func (s *baseSuite) newWorkerWithDB(c *gc.C, db TrackedDB) worker.Worker {
	cfg := WorkerConfig{
		NodeManager: s.nodeManager,
		Clock:       s.clock,
		Logger:      s.logger,
		NewApp: func(driverName string) (DBApp, error) {
			return s.dbApp, nil
		},
		NewDBReplWorker: func(context.Context, DBApp, string, ...TrackedDBWorkerOption) (TrackedDB, error) {
			if s.newDBReplWorker != nil {
				return s.newDBReplWorker()
			}
			return db, nil
		},
	}

	w, err := NewWorker(cfg)
	c.Assert(err, jc.ErrorIsNil)
	return w
}

type dbBaseSuite struct {
	domaintesting.ControllerSuite
	baseSuite
}

func (s *dbBaseSuite) SetUpTest(c *gc.C) {
	s.ControllerSuite.SetUpTest(c)
	s.baseSuite.SetUpTest(c)
}

func (s *dbBaseSuite) TearDownTest(c *gc.C) {
	s.ControllerSuite.TearDownTest(c)
	s.baseSuite.TearDownTest(c)
}

func ensureStartup(c *gc.C, w *dbReplWorker) {
	select {
	case <-w.dbReplReady:
	case <-time.After(jujutesting.LongWait):
		c.Fatal("timed out waiting for Dqlite node start")
	}
}
