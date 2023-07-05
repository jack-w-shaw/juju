// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package modelmigration

import (
	"context"

	"github.com/juju/description/v4"
	"github.com/juju/errors"

	"github.com/juju/juju/core/modelmigration"
	"github.com/juju/juju/domain"
	"github.com/juju/juju/domain/externalcontroller"
	"github.com/juju/juju/domain/externalcontroller/service"
	"github.com/juju/juju/domain/externalcontroller/state"
)

// Coordinator is the interface that is used to add operations to a migration.
type Coordinator interface {
	// Add adds the given operation to the migration.
	Add(modelmigration.Operation)
}

// RegisterImport registers the import operations with the given coordinator.
func RegisterImport(coordinator Coordinator) {
	coordinator.Add(&importOperation{})
}

type ImportService interface {
	ImportExternalControllers(context.Context, []externalcontroller.MigrationControllerInfo) error
}

type importOperation struct {
	modelmigration.BaseOperation

	service ImportService
}

func (i importOperation) Setup(scope modelmigration.Scope) error {
	// We must not use a watcher during migration, so it's safe to pass a
	// nil watcher factory.
	i.service = service.NewService(
		state.NewState(domain.ConstFactory(scope.ControllerDB())), nil)
	return nil
}

// Execute the import on the external controllers description, carefully
// modelling the dependencies we have. External controllers are then inserted
// into the database.
func (i importOperation) Execute(ctx context.Context, model description.Model) error {
	externalControllers := model.ExternalControllers()
	if len(externalControllers) == 0 {
		return nil
	}

	var controllers []externalcontroller.MigrationControllerInfo
	for _, entity := range externalControllers {
		controllers = append(controllers, externalcontroller.MigrationControllerInfo{
			ControllerTag: entity.ID(),
			Alias:         entity.Alias(),
			CACert:        entity.CACert(),
			Addrs:         entity.Addrs(),
			ModelUUIDs:    entity.Models(),
		})
	}

	err := i.service.ImportExternalControllers(ctx, controllers)
	return errors.Annotatef(err, "cannot import external controllers")
}
