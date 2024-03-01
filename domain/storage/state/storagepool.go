// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"fmt"

	"github.com/canonical/sqlair"
	"github.com/juju/errors"

	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/domain"
	domainstorage "github.com/juju/juju/domain/storage"
	storageerrors "github.com/juju/juju/domain/storage/errors"
	"github.com/juju/juju/internal/uuid"
)

// StoragePoolState represents database interactions dealing with storage pools.
type StoragePoolState struct {
	*domain.StateBase
}

type poolAttributes map[string]string

// CreateStoragePool creates a storage pool, returning an error satisfying [storageerrors.PoolAlreadyExists]
// if a pool with the same name already exists.
func (st StoragePoolState) CreateStoragePool(ctx context.Context, pool domainstorage.StoragePoolDetails) error {
	db, err := st.DB()
	if err != nil {
		return errors.Trace(err)
	}

	return CreateStoragePools(ctx, db, []domainstorage.StoragePoolDetails{pool})
}

// CreateStoragePools creates the specified storage pools.
// It is exported for us in the storage/bootstrap package.
func CreateStoragePools(ctx context.Context, db coredatabase.TxnRunner, pools []domainstorage.StoragePoolDetails) error {
	selectUUIDStmt, err := sqlair.Prepare("SELECT &StoragePool.uuid FROM storage_pool WHERE name = $StoragePool.name", StoragePool{})
	if err != nil {
		return errors.Trace(domain.CoerceError(err))
	}

	storagePoolUpserter, err := storagePoolUpserter()
	if err != nil {
		return errors.Trace(domain.CoerceError(err))
	}
	poolAttributesUpdater, err := poolAttributesUpdater()
	if err != nil {
		return errors.Trace(domain.CoerceError(err))
	}

	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		for _, pool := range pools {
			dbPool := StoragePool{Name: pool.Name}
			err := tx.Query(ctx, selectUUIDStmt, dbPool).Get(&dbPool)
			if err != nil && !errors.Is(err, sqlair.ErrNoRows) {
				return errors.Trace(domain.CoerceError(err))
			}
			if err == nil {
				return fmt.Errorf("storage pool %q %w", pool.Name, storageerrors.PoolAlreadyExists)
			}
			poolUUID := uuid.MustNewUUID().String()

			if err := storagePoolUpserter(ctx, tx, poolUUID, pool.Name, pool.Provider); err != nil {
				return errors.Annotatef(err, "creating storage pool %q", pool.Name)
			}

			if err := poolAttributesUpdater(ctx, tx, poolUUID, pool.Attrs); err != nil {
				return errors.Annotatef(err, "creating storage pool %q attributes", pool.Name)
			}
		}
		return nil
	})

	return errors.Trace(domain.CoerceError(err))
}

type upsertStoragePoolFunc func(ctx context.Context, tx *sqlair.TX, poolUUID string, name string, providerType string) error

func storagePoolUpserter() (upsertStoragePoolFunc, error) {
	insertQuery := `
INSERT INTO storage_pool (uuid, name, type)
VALUES (
    $StoragePool.uuid,
    $StoragePool.name,
    $StoragePool.type
)
ON CONFLICT(uuid) DO UPDATE SET name=excluded.name,
                                type=excluded.type
`

	insertStmt, err := sqlair.Prepare(insertQuery, StoragePool{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return func(ctx context.Context, tx *sqlair.TX, poolUUID string, name string, providerType string) error {
		if name == "" {
			return fmt.Errorf("storage pool name cannot be empty%w", errors.Hide(storageerrors.MissingPoolNameError))
		}
		if providerType == "" {
			return fmt.Errorf("storage pool type cannot be empty%w", errors.Hide(storageerrors.MissingPoolTypeError))
		}

		dbPool := StoragePool{
			ID:           poolUUID,
			Name:         name,
			ProviderType: providerType,
		}

		err = tx.Query(ctx, insertStmt, dbPool).Run()
		if err != nil {
			return errors.Trace(err)
		}
		return nil
	}, nil
}

type updatePoolAttributesFunc func(ctx context.Context, tx *sqlair.TX, storagePoolUUID string, attr domainstorage.Attrs) error

func poolAttributesUpdater() (updatePoolAttributesFunc, error) {
	// Delete any keys no longer in the attributes map.
	// TODO(wallyworld) - use sqlair NOT IN operation
	deleteQuery := fmt.Sprintf(`
DELETE FROM  storage_pool_attribute
WHERE        storage_pool_uuid = $M.uuid
-- AND          key NOT IN (?)
`)

	deleteStmt, err := sqlair.Prepare(deleteQuery, sqlair.M{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	insertQuery := `
INSERT INTO storage_pool_attribute
VALUES (
    $poolAttribute.storage_pool_uuid,
    $poolAttribute.key,
    $poolAttribute.value
)
ON CONFLICT(storage_pool_uuid, key) DO UPDATE SET key=excluded.key,
                                                      value=excluded.value
`
	insertStmt, err := sqlair.Prepare(insertQuery, poolAttribute{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return func(ctx context.Context, tx *sqlair.TX, storagePoolUUID string, attr domainstorage.Attrs) error {
		if err := tx.Query(ctx, deleteStmt, sqlair.M{"uuid": storagePoolUUID}).Run(); err != nil {
			return errors.Trace(err)
		}
		for key, value := range attr {
			if err := tx.Query(ctx, insertStmt, poolAttribute{
				ID:    storagePoolUUID,
				Key:   key,
				Value: value,
			}).Run(); err != nil {
				return errors.Trace(err)
			}
		}
		return nil
	}, nil
}

// DeleteStoragePool deletes a storage pool, returning an error satisfying
// [errors.NotFound] if it doesn't exist.
func (st StoragePoolState) DeleteStoragePool(ctx context.Context, name string) error {
	db, err := st.DB()
	if err != nil {
		return errors.Trace(err)
	}

	poolAttributeDeleteQ := `
DELETE FROM storage_pool_attribute
WHERE  storage_pool_attribute.storage_pool_uuid = (select uuid FROM storage_pool WHERE name = $M.name)
`

	poolDeleteQ := `
DELETE FROM storage_pool
WHERE  storage_pool.uuid = (select uuid FROM storage_pool WHERE name = $M.name)
`

	poolAttributeDeleteStmt, err := sqlair.Prepare(poolAttributeDeleteQ, sqlair.M{})
	if err != nil {
		return errors.Trace(err)
	}
	poolDeleteStmt, err := sqlair.Prepare(poolDeleteQ, sqlair.M{})
	if err != nil {
		return errors.Trace(err)
	}

	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		nameMap := sqlair.M{"name": name}
		if err := tx.Query(ctx, poolAttributeDeleteStmt, nameMap).Run(); err != nil {
			return errors.Annotate(err, "deleting storage pool attributes")
		}
		var outcome = sqlair.Outcome{}
		err = tx.Query(ctx, poolDeleteStmt, nameMap).Get(&outcome)
		if err != nil {
			return errors.Trace(err)
		}
		rowsAffected, err := outcome.Result().RowsAffected()
		if err != nil {
			return errors.Annotate(err, "deleting storage pool")
		}
		if rowsAffected == 0 {
			return fmt.Errorf("storage pool %q not found%w", name, errors.Hide(storageerrors.PoolNotFoundError))
		}
		return errors.Annotate(err, "deleting storage pool")
	})
	return errors.Trace(domain.CoerceError(err))
}

// ReplaceStoragePool replaces an existing storage pool, returning an error
// satisfying [errors.NotFound] if a pool with the name does not exist.
func (st StoragePoolState) ReplaceStoragePool(ctx context.Context, pool domainstorage.StoragePoolDetails) error {
	db, err := st.DB()
	if err != nil {
		return errors.Trace(err)
	}

	selectUUIDStmt, err := sqlair.Prepare("SELECT &StoragePool.uuid FROM storage_pool WHERE name = $StoragePool.name", StoragePool{})
	if err != nil {
		return errors.Trace(domain.CoerceError(err))
	}

	storagePoolUpserter, err := storagePoolUpserter()
	if err != nil {
		return errors.Trace(domain.CoerceError(err))
	}
	poolAttributesUpdater, err := poolAttributesUpdater()
	if err != nil {
		return errors.Trace(domain.CoerceError(err))
	}

	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		dbPool := StoragePool{Name: pool.Name}
		err := tx.Query(ctx, selectUUIDStmt, dbPool).Get(&dbPool)
		if err != nil && !errors.Is(err, sqlair.ErrNoRows) {
			return errors.Trace(err)
		}
		if err != nil {
			return fmt.Errorf("storage pool %q not found%w", pool.Name, errors.Hide(storageerrors.PoolNotFoundError))
		}
		poolUUID := dbPool.ID
		if err := storagePoolUpserter(ctx, tx, poolUUID, pool.Name, pool.Provider); err != nil {
			return errors.Annotate(err, "updating storage pool")
		}

		if err := poolAttributesUpdater(ctx, tx, poolUUID, pool.Attrs); err != nil {
			return errors.Annotatef(err, "updating storage pool %s attributes", poolUUID)
		}
		return nil
	})

	return errors.Trace(domain.CoerceError(err))
}

type loadStoragePoolsFunc func(ctx context.Context, tx *sqlair.TX) ([]domainstorage.StoragePoolDetails, error)

func (st StoragePoolState) storagePoolsLoader(filter domainstorage.StoragePoolFilter) (loadStoragePoolsFunc, error) {
	query := `
SELECT (sp.uuid, sp.name, sp.type) AS (&StoragePool.*),
       (sp_attr.key, sp_attr.value) AS (&poolAttribute.*)
FROM   storage_pool sp
       LEFT JOIN storage_pool_attribute sp_attr ON sp_attr.storage_pool_uuid = sp.uuid
`

	types := []any{
		StoragePool{},
		poolAttribute{},
	}

	var queryArgs []any
	condition, args := st.buildFilter(filter)
	if len(args) > 0 {
		query = query + "WHERE " + condition
		types = append(types, args...)
		queryArgs = append([]any{}, args...)
	}

	queryStmt, err := sqlair.Prepare(query, types...)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return func(ctx context.Context, tx *sqlair.TX) ([]domainstorage.StoragePoolDetails, error) {
		var (
			dbRows    StoragePools
			keyValues []poolAttribute
		)
		err = tx.Query(ctx, queryStmt, queryArgs...).GetAll(&dbRows, &keyValues)
		if err != nil {
			return nil, errors.Annotate(err, "loading storage pool")
		}
		return dbRows.toStoragePools(keyValues)
	}, nil
}

// ListStoragePools returns the storage pools matching the specified filter.
func (st StoragePoolState) ListStoragePools(ctx context.Context, filter domainstorage.StoragePoolFilter) ([]domainstorage.StoragePoolDetails, error) {
	db, err := st.DB()
	if err != nil {
		return nil, errors.Trace(err)
	}

	storagePoolsLoader, err := st.storagePoolsLoader(filter)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var result []domainstorage.StoragePoolDetails
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		var err error
		result, err = storagePoolsLoader(ctx, tx)
		return errors.Trace(err)
	})
	return result, errors.Trace(domain.CoerceError(err))
}

func (st StoragePoolState) buildFilter(filter domainstorage.StoragePoolFilter) (string, []any) {
	if len(filter.Names) == 0 && len(filter.Providers) == 0 {
		return "", nil
	}

	if len(filter.Names) > 0 && len(filter.Providers) > 0 {
		condition := "sp.name IN ($StoragePoolNames[:]) AND sp.type IN ($StorageProviderTypes[:])"
		return condition, []any{StoragePoolNames(filter.Names), StorageProviderTypes(filter.Providers)}
	}

	if len(filter.Names) > 0 {
		condition := "sp.name IN ($StoragePoolNames[:])"
		return condition, []any{StoragePoolNames(filter.Names)}
	}

	condition := "sp.type IN ($StorageProviderTypes[:])"
	return condition, []any{StorageProviderTypes(filter.Providers)}
}

// GetStoragePoolByName returns the storage pool with the specified name, returning an error
// satisfying [storageerrors.PoolNotFoundError] if it doesn't exist.
func (st StoragePoolState) GetStoragePoolByName(ctx context.Context, name string) (domainstorage.StoragePoolDetails, error) {
	db, err := st.DB()
	if err != nil {
		return domainstorage.StoragePoolDetails{}, errors.Trace(err)
	}

	storagePoolsLoader, err := st.storagePoolsLoader(domainstorage.StoragePoolFilter{
		Names: []string{name},
	})
	if err != nil {
		return domainstorage.StoragePoolDetails{}, errors.Trace(err)
	}

	var storagePools []domainstorage.StoragePoolDetails
	err = db.Txn(ctx, func(ctx context.Context, tx *sqlair.TX) error {
		var err error
		storagePools, err = storagePoolsLoader(ctx, tx)
		return errors.Trace(err)
	})
	if err != nil {
		return domainstorage.StoragePoolDetails{}, errors.Trace(err)
	}
	if len(storagePools) == 0 {
		return domainstorage.StoragePoolDetails{}, fmt.Errorf("storage pool %q %w", name, storageerrors.PoolNotFoundError)
	}
	if len(storagePools) > 1 {
		return domainstorage.StoragePoolDetails{}, errors.Errorf("expected 1 storage pool, got %d", len(storagePools))
	}
	return storagePools[0], errors.Trace(domain.CoerceError(err))
}
