// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"time"

	"github.com/juju/version/v2"
)

type Info struct {
	PreviousVersion  version.Number
	TargetVersion    version.Number
	CreatedAt        time.Time
	StartedAt        time.Time
	CompletedAt      time.Time
	ControllersReady []string
	ControllersDone  []string
}
