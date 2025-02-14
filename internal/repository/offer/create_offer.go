package offer

import (
	"context"

	"teklif/internal/models/domain"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (r *Repo) CreateOffer(ctx context.Context, offer domain.Offer) error {
	query := `
		INSERT INTO special_offers (
			title,
			description,
			type,
			expiry_date,
			commission_agent,
			commission_agency,
			discount,
			price_installment,
			price_cash,
			status,
			media
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.pool.Exec(ctx, query,
		offer.Title, offer.Description, offer.Type, offer.ExpiryDate,
		offer.CommissionAgent, offer.CommissionAgency, offer.Discount,
		offer.PriceInstallment, offer.PriceCash, offer.Status,
		pq.Array(offer.Media),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
