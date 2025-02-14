package offer

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repo) ExpireOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) error {
	query := `
		UPDATE special_offers
		SET status = 'expired'
		WHERE id IN (
			SELECT offer_id 
			FROM offer_properties 
			WHERE property_id = ANY($1)
		)
		AND status = 'approved';
	`

	_, err := r.pool.Exec(ctx, query, propertyIDs)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
