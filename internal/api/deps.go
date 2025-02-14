package api

import (
	"context"
	"teklif/internal/models/domain"

	"github.com/google/uuid"
)

type offerService interface {
	CreateOffer(ctx context.Context, offer domain.Offer) error
	ApproveOffer(ctx context.Context, offerID uuid.UUID) error
	AttachOffer(ctx context.Context, attachOffer domain.AttachOffer) error
	GetActiveOffers(ctx context.Context) ([]domain.Offer, error)
	GetActiveOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) ([]domain.Offer, error)
	GetWeekExpiredOffers(ctx context.Context) ([]domain.Offer, error)
	ExpireOffersByProperties(ctx context.Context, propertyIDs []uuid.UUID) error
	DeleteOffer(ctx context.Context, offerID uuid.UUID) error
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Error(text ...any)
}
