package offer

import (
	"context"
	"teklif/internal/models/domain"

	"github.com/google/uuid"
)

func (s *Service) GetActiveOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) ([]domain.Offer, error) {
	return s.offerRepo.GetActiveOffersByProperties(ctx, propertyIDs)
}