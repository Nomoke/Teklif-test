package offer

import (
	"context"
	"teklif/internal/api/models"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repo) ExpireOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) error {
	query := `
		UPDATE special_offers
		SET status = $1
		WHERE id IN (
			SELECT offer_id 
			FROM offer_properties 
			WHERE property_id = ANY($2)
		)
		AND status = $3;
	`

	_, err := r.pool.Exec(ctx, query, models.OfferStatusExpired, propertyIDs, models.OfferStatusApproved)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
