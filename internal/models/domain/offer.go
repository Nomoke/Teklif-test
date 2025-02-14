package domain

import (
	"time"

	"github.com/google/uuid"
)

type Offer struct {
	ID               uuid.UUID
	Title            string
	Description      string
	Type             string
	ExpiryDate       time.Time
	CommissionAgent  *float64
	CommissionAgency *float64
	Discount         *float64
	PriceInstallment *float64
	PriceCash        *float64
	Status           string
	Media            []string
}

type AttachOffer struct {
	OfferID     uuid.UUID
	PropertyIDs []uuid.UUID
}
