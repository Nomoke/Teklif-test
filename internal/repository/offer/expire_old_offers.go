package offer

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Repo) ExpireOldOffers(ctx context.Context) error {
	query := `
		UPDATE special_offers 
		SET status = 'expired'
		WHERE expiry_date <= NOW() AND status != 'expired'
	`

	_, err := s.pool.Exec(ctx, query)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
