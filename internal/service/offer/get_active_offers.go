package offer

import (
	"context"
	"teklif/internal/models/domain"
)

func (s *Service) GetActiveOffers(ctx context.Context) ([]domain.Offer, error) {
	return s.offerRepo.GetActiveOffers(ctx)
}