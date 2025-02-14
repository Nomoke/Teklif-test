package offer

import (
	"context"
	"teklif/internal/models/domain"

	"github.com/google/uuid"
)

type offerRepo interface {
	CreateOffer(ctx context.Context, offer domain.Offer) error
	ApproveOffer(ctx context.Context, offerID uuid.UUID) error
	AttachOffer(ctx context.Context, attachOffer domain.AttachOffer) error
	GetActiveOffers(ctx context.Context) ([]domain.Offer, error)
	GetActiveOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) ([]domain.Offer, error)
	GetWeekExpiredOffers(ctx context.Context) ([]domain.Offer, error)
	ExpireOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) error
	DeleteOffer(ctx context.Context, offerID uuid.UUID) error
	ExpireOldOffers(ctx context.Context) error
}
