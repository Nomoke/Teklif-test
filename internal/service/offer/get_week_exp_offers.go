package offer

import (
	"context"
	"teklif/internal/models/domain"
)

func (s *Service) GetWeekExpiredOffers(ctx context.Context) ([]domain.Offer, error) {
	return s.offerRepo.GetWeekExpiredOffers(ctx)
}
