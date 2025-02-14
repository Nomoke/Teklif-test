package offer

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repo) ApproveOffer(ctx context.Context, offerID uuid.UUID) error {
	query := `
		UPDATE special_offers
		SET status = 'approved', updated_at = NOW()
		WHERE id = $1 AND status = 'pending'
	`

	_, err := r.pool.Exec(ctx, query, offerID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
