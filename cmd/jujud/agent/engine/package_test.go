// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package engine_test

import (
	"testing"

	"github.com/juju/juju/api/base"
	"github.com/juju/worker/v3"
	gc "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}

type dummyWorker struct {
	worker.Worker
}

type dummyAPICaller struct {
	base.APICaller
}
