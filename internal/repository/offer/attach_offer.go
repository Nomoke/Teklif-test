package offer

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"teklif/internal/models/domain"
)

func (r *Repo) AttachOffer(ctx context.Context, aOffer domain.AttachOffer) error {
	query := `
		INSERT INTO offer_properties (offer_id, property_id) 
		VALUES ($1, $2) ON CONFLICT DO NOTHING;
	`

	batch := &pgx.Batch{}

	for _, propertyID := range aOffer.PropertyIDs {
		batch.Queue(query, aOffer.OfferID, propertyID)
	}

	br := r.pool.SendBatch(ctx, batch)
	defer br.Close()

	for range aOffer.PropertyIDs {
		if _, err := br.Exec(); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
