package offer

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) ApproveOffer(ctx context.Context, offerID uuid.UUID) error {
	return s.offerRepo.ApproveOffer(ctx, offerID)
}
