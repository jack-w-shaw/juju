// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"

	"github.com/canonical/sqlair"
	"github.com/juju/collections/set"
	"github.com/juju/collections/transform"
	"github.com/juju/errors"
	"github.com/juju/utils/v3"
	"github.com/juju/version/v2"

	"github.com/juju/juju/domain"
)

// State is used to access the database.
type State struct {
	*domain.StateBase
}

var noUpgradeErr = errors.Errorf("no current upgrade")

// NewState creates a state to access the database.
func NewState(factory domain.DBFactory) *State {
	return &State{
		StateBase: domain.NewStateBase(factory),
	}
}

// EnsureUpgradeInfo returns the current upgrade info and ensures the
// current controller is included and made ready. If no upgrade exists,
// one is created
func (st *State) EnsureUpgradeInfo(ctx context.Context, controllerID string, previousVersion, targetVersion version.Number) (Info, []InfoControllerNode, error) {
	db, err := st.DB()
	if err != nil {
		return Info{}, nil, errors.Trace(err)
	}

	var (
		resInfo      Info
		resNodeInfos []InfoControllerNode
	)
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		var err error
		resInfo, resNodeInfos, err = st.getCurrentUpgrade(ctx, tx)
		if err != nil && err != noUpgradeErr {
			return errors.Trace(err)
		}
		if err == noUpgradeErr {
			resInfo, resNodeInfos, err = st.initialiseUpgrade(ctx, tx, controllerID, previousVersion, targetVersion)
			return errors.Trace(err)
		}
		if err := verifyVersionMatches(resInfo, previousVersion, targetVersion); err != nil {
			return errors.Trace(err)
		}
		resInfo, resNodeInfos, err = st.ensureNodeReady(ctx, tx, resInfo, resNodeInfos, controllerID)
		if err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(err)
	})
	return resInfo, resNodeInfos, errors.Trace(err)
}

func (st *State) getCurrentUpgrade(ctx context.Context, tx *sqlair.TX) (Info, []InfoControllerNode, error) {
	q1 := `
SELECT (uuid, previous_version, target_version, created_at, started_at) AS &Info.* FROM upgrade_info AS info
WHERE info.completed_at IS NULL
ORDER BY info.created_at DESC LIMIT 1`
	s1, err := sqlair.Prepare(q1, Info{})
	if err != nil {
		return Info{}, nil, errors.Annotatef(err, "preparing %q", q1)
	}

	q2 := `
SELECT (controller_node_id, status) AS &InfoControllerNode.*
FROM   upgrade_info_controller_node
       JOIN upgrade_node_status AS status
       ON upgrade_info_controller_node.upgrade_node_status_id = status.id
WHERE upgrade_info_controller_node.upgrade_info_uuid = $M.info_uuid `
	s2, err := sqlair.Prepare(q2, InfoControllerNode{}, sqlair.M{})
	if err != nil {
		return Info{}, nil, errors.Annotatef(err, "preparing %q", q2)
	}

	var info Info
	err = tx.Query(ctx, s1).Get(&info)
	if err == sqlair.ErrNoRows {
		return Info{}, nil, noUpgradeErr
	}
	if err != nil {
		return Info{}, nil, errors.Trace(err)
	}

	var nodeInfos []InfoControllerNode
	err = tx.Query(ctx, s2, sqlair.M{"info_uuid": info.UUID}).GetAll(&nodeInfos)
	if err != nil && err != sqlair.ErrNoRows {
		return Info{}, nil, errors.Trace(err)
	}
	return info, nodeInfos, nil
}

func (st *State) initialiseUpgrade(
	ctx context.Context,
	tx *sqlair.TX,
	controllerID string,
	previousVersion version.Number,
	targetVersion version.Number,
) (Info, []InfoControllerNode, error) {
	q := `
INSERT INTO upgrade_info (uuid, previous_version, target_version, created_at) VALUES
($M.uuid, $M.previousVersion, $M.targetVersion, DATETIME('now'))`
	s, err := sqlair.Prepare(q, sqlair.M{})
	if err != nil {
		return Info{}, nil, errors.Annotatef(err, "preparing %q", q)
	}
	err = tx.Query(ctx, s, sqlair.M{
		"uuid":            utils.MustNewUUID().String(),
		"previousVersion": previousVersion.String(),
		"targetVersion":   targetVersion.String(),
	}).Run()
	if err != nil {
		return Info{}, nil, errors.Trace(err)
	}

	info, nodeInfos, err := st.getCurrentUpgrade(ctx, tx)
	if err != nil {
		return Info{}, nil, errors.Trace(err)
	}
	info, nodeInfos, err = st.ensureNodeReady(ctx, tx, info, nodeInfos, controllerID)
	if err != nil {
		return Info{}, nil, errors.Trace(err)
	}
	return info, nodeInfos, nil
}

func (st *State) ensureNodeReady(
	ctx context.Context,
	tx *sqlair.TX,
	info Info,
	nodeInfos []InfoControllerNode,
	controllerID string,
) (Info, []InfoControllerNode, error) {
	for _, nodeInfo := range nodeInfos {
		if nodeInfo.ControllerNodeID == controllerID {
			return info, nodeInfos, nil
		}
	}
	q := `
INSERT INTO upgrade_info_controller_node (uuid, controller_node_id, upgrade_info_uuid, upgrade_node_status_id) VALUES
($M.uuid, $M.controllerID, $M.infoUUID, $M.readyKey)`
	s, err := sqlair.Prepare(q, sqlair.M{})
	if err != nil {
		return Info{}, nil, errors.Annotatef(err, "preparing %q", q)
	}
	err = tx.Query(ctx, s, sqlair.M{
		"uuid":         utils.MustNewUUID().String(),
		"controllerID": controllerID,
		"infoUUID":     info.UUID,
		"readyKey":     0,
	}).Run()
	if err != nil {
		return Info{}, nil, errors.Trace(err)
	}
	nodeInfos = append(nodeInfos, InfoControllerNode{ControllerNodeID: controllerID, NodeStatus: "ready"})
	return info, nodeInfos, nil
}

func verifyVersionMatches(info Info, previousVersion, targetVersion version.Number) error {
	if info.PreviousVersion != previousVersion.String() || info.TargetVersion != targetVersion.String() {
		return errors.NotValidf(
			"current upgrade (%s -> %s) does not match started upgrade (%s -> %s)",
			info.PreviousVersion, info.TargetVersion, previousVersion, targetVersion,
		)
	}
	return nil
}

// AllProvisionedControllersReady returns true if and only if all controllers
// that have been started by the provisioner have called EnsureUpgradeInfo with
// matching versions.
func (st *State) AllProvisionedControllersReady(ctx context.Context) (bool, error) {
	db, err := st.DB()
	if err != nil {
		return false, errors.Trace(err)
	}
	var (
		provisionedControllers []string
		nodeInfos              []InfoControllerNode
	)
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		var err error
		provisionedControllers, err = st.getProvisionedControllers(ctx, tx)
		if err != nil {
			return errors.Trace(err)
		}
		_, nodeInfos, err = st.getCurrentUpgrade(ctx, tx)
		if err != nil {
			return errors.Trace(err)
		}
		return nil
	})
	ready := set.NewStrings(transform.Slice(nodeInfos, func(info InfoControllerNode) string { return info.ControllerNodeID })...)
	missing := set.NewStrings(provisionedControllers...).Difference(ready)
	return missing.IsEmpty(), nil
}

func (st *State) getProvisionedControllers(ctx context.Context, tx *sqlair.TX) ([]string, error) {
	q := `SELECT controller_id FROM controller_node WHERE dqlite_node_id IS NOT NULL`
	s, err := sqlair.Prepare(q)
	if err != nil {
		return nil, errors.Annotatef(err, "preparing %q", q)
	}
	var controllers []string
	err = tx.Query(ctx, s).GetAll(&controllers)
	return controllers, errors.Trace(err)
}

// StartUpgrade starts the current upgrade if it exists
func (st *State) StartUpgrade(ctx context.Context) error {
	db, err := st.DB()
	if err != nil {
		return errors.Trace(err)
	}
	return errors.Trace(db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		info, _, err := st.getCurrentUpgrade(ctx, tx)
		if err != nil {
			return errors.Trace(err)
		}
		if info.StartedAt.Valid {
			return nil
		}
		q := `
UPDATE upgrade_info
SET started_at = DATETIME("now")
WHERE uuid = $M.uuid`
		s, err := sqlair.Prepare(q, sqlair.M{})
		if err != nil {
			return errors.Annotatef(err, "prearing %q", q)
		}
		return tx.Query(ctx, s, sqlair.M{"uuid": info.UUID}).Run()
	}))
}

func (st *State) SetControllerDone(ctx context.Context, controllerID string) error {
	db, err := st.DB()
	if err != nil {
		return errors.Trace(err)
	}
	return errors.Trace(db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		info, nodeInfos, err := st.getCurrentUpgrade(ctx, tx)
		if err != nil {
			return errors.Trace(err)
		}
		for _, nodeInfo := range nodeInfos {
			if nodeInfo.ControllerNodeID == controllerID {
				switch nodeInfo.NodeStatus {
				case "done":
					return nil
				case "ready":
					err := st.setNodeToDone(ctx, tx, info.UUID, controllerID)
					if err != nil {
						return errors.Trace(err)
					}
					err = st.maybeCompleteUpgrade(ctx, tx)
					if err != nil {
						return errors.Trace(err)
					}

				}
			}
		}
		return errors.NotFoundf("controller %q", controllerID)
	}))
}

func (st *State) setNodeToDone(ctx context.Context, tx *sqlair.TX, infoUUID string, controllerID string) error {
	q := `
UPDATE upgrade_info_controller_node
SET upgrade_node_status_id = $M.done_key
WHERE upgrade_info_uuid = $M.info_uuid AND controller_node_id = $M.controller_id`
	s, err := sqlair.Prepare(q, sqlair.M{})
	if err != nil {
		return errors.Annotatef(err, "preparing %q", q)
	}
	return errors.Trace(tx.Query(ctx, s, sqlair.M{
		"done_key":      "1",
		"info_uuid":     infoUUID,
		"controller_id": controllerID,
	}).Run())
}

func (st *State) maybeCompleteUpgrade(ctx context.Context, tx *sqlair.TX) error {
	info, nodeInfos, err := st.getCurrentUpgrade(ctx, tx)
	if err != nil {
		return errors.Trace(err)
	}
	for _, nodeInfo := range nodeInfos {
		if nodeInfo.NodeStatus != "done" {
			return nil
		}
	}
	q := `
UPDATE upgrade_info
SET completion_time = DATETIME("now")
WHERE uuid = $M.info_uuid`
	s, err := sqlair.Prepare(q, sqlair.M{})
	if err != nil {
		return errors.Annotatef(err, "prearing %q", q)
	}
	err = tx.Query(ctx, s, sqlair.M{
		"info_uuid": info.UUID,
	}).Run()
	return errors.Trace(err)
}

// IsUpgrading returns true if an upgrade is currently in progress.
func (st *State) IsUpgrading(ctx context.Context) (bool, error) {
	db, err := st.DB()
	if err != nil {
		return false, errors.Trace(err)
	}
	var upgrading bool
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		_, _, err := st.getCurrentUpgrade(ctx, tx)
		if err == nil {
			upgrading = true
			return nil
		}
		if err == noUpgradeErr {
			upgrading = false
			return nil
		}
		return errors.Trace(err)
	})
	return upgrading, errors.Trace(err)
}
