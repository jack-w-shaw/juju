// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	"github.com/juju/errors"

	"github.com/juju/juju/cloud"
	"github.com/juju/juju/domain/cloud/state"
)

// State describes retrieval and persistence methods for storage.
type State interface {
	// Save persists the input cloud entity.
	Save(context.Context, cloud.Cloud) error

	// List returns the clouds matching the optional filter.
	List(context.Context, *state.Filter) ([]cloud.Cloud, error)
}

// Service provides the API for working with clouds.
type Service struct {
	st State
}

// NewService returns a new service reference wrapping the input state.
func NewService(st State) *Service {
	return &Service{st}
}

// Save inserts or updates the specified cloud.
func (s *Service) Save(ctx context.Context, cloud cloud.Cloud) error {
	err := s.st.Save(ctx, cloud)
	return errors.Annotate(err, "updating cloud state")
}

// ListAll returns all the clouds.
func (s *Service) ListAll(ctx context.Context) ([]cloud.Cloud, error) {
	all, err := s.st.List(ctx, nil)
	return all, errors.Trace(err)
}

// Get returns the named cloud.
func (s *Service) Get(ctx context.Context, name string) (*cloud.Cloud, error) {
	clouds, err := s.st.List(ctx, &state.Filter{Name: name})
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(clouds) == 0 {
		return nil, errors.NotFoundf("cloud %q", name)
	}
	result := clouds[0]
	return &result, nil
}
