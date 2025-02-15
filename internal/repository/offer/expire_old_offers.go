package offer

import (
	"context"

	"teklif/internal/api/models"

	"github.com/pkg/errors"
)

func (s *Repo) ExpireOldOffers(ctx context.Context) error {
	query := `
		UPDATE special_offers 
		SET status = $1
		WHERE expiry_date <= NOW() AND status != $1
	`

	_, err := s.pool.Exec(ctx, query, models.OfferStatusExpired)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
