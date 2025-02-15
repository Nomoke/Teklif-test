package offer

import (
	"context"
	"teklif/internal/api/models"
	"teklif/internal/models/domain"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *Repo) GetActiveOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) ([]domain.Offer, error) {
	query := `
		SELECT
			so.id,
			so.title,
			so.description,
			so.type,
			so.expiry_date,
		    so.commission_agent,
			so.commission_agency,
			so.discount,
		    so.price_installment,
			so.price_cash,
			so.status,
			so.media
		FROM special_offers so
		JOIN offer_properties op ON so.id = op.offer_id
		WHERE op.property_id = ANY($1) AND so.status = $2 AND so.expiry_date > NOW();
	`

	rows, err := r.pool.Query(ctx, query, propertyIDs, models.OfferStatusApproved)
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
