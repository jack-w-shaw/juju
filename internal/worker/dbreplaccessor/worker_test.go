// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package dbreplaccessor

import (
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4"
	"github.com/juju/worker/v4/workertest"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/database"
)

type workerSuite struct {
	baseSuite

	trackedDB *MockTrackedDB
	driver    *MockDriver
}

var _ = gc.Suite(&workerSuite{})

func (s *workerSuite) TestKilledGetDBErrDying(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(c.MkDir(), nil)
	mgrExp.DqliteSQLDriver(gomock.Any()).Return(s.driver, nil)

	w := s.newWorker(c)
	defer workertest.DirtyKill(c, w)

	dbw := w.(*dbReplWorker)
	ensureStartup(c, dbw)

	w.Kill()

	_, err := dbw.GetDB("anything")
	c.Assert(err, jc.ErrorIs, database.ErrDBReplAccessorDying)
}

func (s *workerSuite) TestGetDB(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(c.MkDir(), nil)
	mgrExp.DqliteSQLDriver(gomock.Any()).Return(s.driver, nil)

	done := make(chan struct{})

	s.expectTrackedDB(done)

	w := s.newWorker(c)
	defer func() {
		close(done)
		workertest.DirtyKill(c, w)
	}()

	dbw := w.(*dbReplWorker)
	ensureStartup(c, dbw)

	runner, err := dbw.GetDB("anything")
	c.Assert(err, jc.ErrorIsNil)
	c.Check(runner, gc.NotNil)
}

func (s *workerSuite) TestGetDBNotFound(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.newDBReplWorker = func() (TrackedDB, error) {
		return nil, database.ErrDBNotFound
	}

	s.expectClock()

	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(c.MkDir(), nil)
	mgrExp.DqliteSQLDriver(gomock.Any()).Return(s.driver, nil)

	w := s.newWorker(c)
	defer workertest.DirtyKill(c, w)

	dbw := w.(*dbReplWorker)
	ensureStartup(c, dbw)

	_, err := dbw.GetDB("other")

	// The error isn't passed through, although we really should expose this
	// in the runner.
	c.Assert(err, jc.ErrorIs, worker.ErrDead)
}

func (s *workerSuite) TestGetDBFound(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClock()

	mgrExp := s.nodeManager.EXPECT()
	mgrExp.EnsureDataDir().Return(c.MkDir(), nil)
	mgrExp.DqliteSQLDriver(gomock.Any()).Return(s.driver, nil)

	done := make(chan struct{})

	s.expectTrackedDB(done)

	w := s.newWorker(c)
	defer func() {
		close(done)
		workertest.DirtyKill(c, w)
	}()

	dbw := w.(*dbReplWorker)
	ensureStartup(c, dbw)

	runner, err := dbw.GetDB("anything")
	c.Assert(err, jc.ErrorIsNil)
	c.Check(runner, gc.NotNil)

	// Notice that no additional changes are expected.

	runner, err = dbw.GetDB("anything")
	c.Assert(err, jc.ErrorIsNil)
	c.Check(runner, gc.NotNil)
}

func (s *workerSuite) newWorker(c *gc.C) worker.Worker {
	return s.newWorkerWithDB(c, s.trackedDB)
}

func (s *workerSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := s.baseSuite.setupMocks(c)

	s.trackedDB = NewMockTrackedDB(ctrl)
	s.driver = NewMockDriver(ctrl)

	return ctrl
}

func (s *workerSuite) expectTrackedDB(done chan struct{}) {
	s.trackedDB.EXPECT().Kill().AnyTimes()
	s.trackedDB.EXPECT().Wait().DoAndReturn(func() error {
		<-done
		return nil
	})
}
