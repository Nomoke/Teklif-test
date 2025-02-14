package offer

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteOffer(ctx context.Context, offerID uuid.UUID) error {
	return s.offerRepo.DeleteOffer(ctx, offerID)
}
