package offer

import (
	"context"
	"teklif/internal/models/domain"
)

func (s *Service) CreateOffer(ctx context.Context, offer domain.Offer) error {
	return s.offerRepo.CreateOffer(ctx, offer)
}
