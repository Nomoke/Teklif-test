package offer

import (
	"context"
	"teklif/internal/api/models"
	"teklif/internal/models/domain"

	"github.com/pkg/errors"
)

func (r *Repo) GetActiveOffers(ctx context.Context) ([]domain.Offer, error) {
	query := `
		SELECT
			id,
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
		FROM special_offers
		WHERE status = $1 AND expiry_date > NOW();
`

	rows, err := r.pool.Query(ctx, query, models.OfferStatusApproved)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer rows.Close()

	var offers []domain.Offer

	for rows.Next() {
		var offer domain.Offer
		if err := rows.Scan(
			&offer.ID,
			&offer.Title,
			&offer.Description,
			&offer.Type,
			&offer.ExpiryDate,
			&offer.CommissionAgent,
			&offer.CommissionAgency,
			&offer.Discount,
			&offer.PriceInstallment,
			&offer.PriceCash,
			&offer.Status,
			&offer.Media,
		); err != nil {
			return nil, errors.WithStack(err)
		}

		offers = append(offers, offer)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.WithStack(err)
	}

	return offers, nil
}
