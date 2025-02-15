package offer

import (
	"context"

	"teklif/internal/api/models"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repo) ApproveOffer(ctx context.Context, offerID uuid.UUID) error {
	query := `
		UPDATE special_offers
		SET status = $1, updated_at = NOW()
		WHERE id = $2 AND status = $3
	`

	_, err := r.pool.Exec(ctx, query, models.OfferStatusApproved, offerID, models.OfferStatusPending)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
