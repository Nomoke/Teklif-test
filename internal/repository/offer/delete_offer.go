package offer

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repo) DeleteOffer(ctx context.Context, offerID uuid.UUID) error {
	query := `
		DELETE FROM special_offers
		WHERE id = $1
	`

	_, err := r.pool.Exec(ctx, query, offerID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
