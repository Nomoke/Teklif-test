package offer

import (
	"context"

	"teklif/internal/models/domain"
)

func (s *Service) AttachOffer(ctx context.Context, attachOffer domain.AttachOffer) error {
	return s.offerRepo.AttachOffer(ctx, attachOffer)
}
