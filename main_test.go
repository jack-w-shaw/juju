package juju

import (
	"context"
	"testing"

	"github.com/canonical/sqlair"
	"github.com/juju/tc"

	schematesting "github.com/juju/juju/domain/schema/testing"
)

type suite struct {
	schematesting.ModelSuite
}

func TestSuite(t *testing.T) {
	tc.Run(t, &suite{})
}

type (
	life struct {
		ID    int    `db:"id"`
		Value string `db:"value"`
	}
	ids []int
)

func (s *suite) TestTestSingular(c *tc.C) {
	stmt, err := sqlair.Prepare(`
SELECT &life.*
FROM life
WHERE id = ($ids[:])
	`, life{}, ids{})
	c.Assert(err, tc.ErrorIsNil)

	lifeIDs := ids([]int{0})
	err = s.TxnRunner().Txn(c.Context(), func(ctx context.Context, tx *sqlair.TX) error {
		return tx.Query(ctx, stmt, lifeIDs).GetAll(&[]life{})
	})
	c.Assert(err, tc.ErrorIsNil)
}

func (s *suite) TestTestMultiple(c *tc.C) {
	stmt, err := sqlair.Prepare(`
SELECT &life.*
FROM life
WHERE id = ($ids[:])
	`, life{}, ids{})
	c.Assert(err, tc.ErrorIsNil)

	lifeIDs := ids([]int{0, 1})
	err = s.TxnRunner().Txn(c.Context(), func(ctx context.Context, tx *sqlair.TX) error {
		return tx.Query(ctx, stmt, lifeIDs).GetAll(&[]life{})
	})
	c.Assert(err, tc.ErrorIsNil)
}
