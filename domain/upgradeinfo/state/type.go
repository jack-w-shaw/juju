// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import "database/sql"

// These structs represent the persistent upgrade info schema in the database.

type Info struct {
	UUID            string         `db:"uuid"`
	PreviousVersion string         `db:"previous_version"`
	TargetVersion   string         `db:"target_version"`
	CreatedAt       string         `db:"created_at"`
	StartedAt       sql.NullString `db:"started_at"`
	CompletedAt     sql.NullString `db:"completed_at"`
}

type InfoControllerNode struct {
	ControllerNodeID string `db:"controller_node_id"`
	NodeStatus       string `db:"status"`
}
