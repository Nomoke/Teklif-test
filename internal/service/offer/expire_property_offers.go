package offer

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) ExpireOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) error {
	return s.offerRepo.ExpireOffersByProperties(ctx, propertyIDs)
}
