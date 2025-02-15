package models

import (
	"time"

	"teklif/internal/models/domain"

	"github.com/google/uuid"
)

type OfferResp struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Type             string    `json:"type"`
	ExpiryDate       time.Time `json:"expiry_date"`
	CommissionAgent  *float64  `json:"commission_agent"`
	CommissionAgency *float64  `json:"commission_agency"`
	Discount         *float64  `json:"discount"`
	PriceInstallment *float64  `json:"price_installment"`
	PriceCash        *float64  `json:"price_cash"`
	Status           string    `json:"status"`
	Media            []string  `json:"media"`
}

type CreateOfferReq struct {
	Title            string    `json:"title" form:"title" binding:"required"`
	Description      string    `json:"description" form:"description" binding:"required"`
	Type             string    `json:"type" form:"type" binding:"required"`
	ExpiryDate       time.Time `json:"expiry_date" form:"expiry_date" binding:"required"`
	CommissionAgent  *float64  `json:"commission_agent" form:"commission_agent"`
	CommissionAgency *float64  `json:"commission_agency" form:"commission_agency"`
	Discount         *float64  `json:"discount" form:"discount"`
	PriceInstallment *float64  `json:"price_installment" form:"price_installment"`
	PriceCash        *float64  `json:"price_cash" form:"price_cash"`
}

const (
	OfferStatusPending  = "pending"
	OfferStatusApproved = "approved"
	OfferStatusExpired  = "expired"
)

type AttachOfferPropertiesReq struct {
	// OfferID UUID
	// @example "123e4567-e89b-12d3-a456-426614174000"
	OfferID uuid.UUID `json:"offer_id" binding:"required"`

	// PropertyIDs list of UUIDs
	// @example ["123e4567-e89b-12d3-a456-426614174001", "123e4567-e89b-12d3-a456-426614174002"]
	PropertyIDs []uuid.UUID `json:"property_ids" binding:"required"`
}

type PropertyIDsReq struct {
	// PropertyIDs list of UUIDs
	// @example ["123e4567-e89b-12d3-a456-426614174001", "123e4567-e89b-12d3-a456-426614174002"]
	PropertyIDs []uuid.UUID `json:"property_ids" binding:"required"`
}

func BuildOffersResponse(offer []domain.Offer) []OfferResp {
	offers := make([]OfferResp, len(offer))

	for i, o := range offer {
		offers[i] = OfferResp{
			ID:               o.ID,
			Title:            o.Title,
			Description:      o.Description,
			Type:             o.Type,
			ExpiryDate:       o.ExpiryDate,
			CommissionAgent:  o.CommissionAgent,
			CommissionAgency: o.CommissionAgency,
			Discount:         o.Discount,
			PriceInstallment: o.PriceInstallment,
			PriceCash:        o.PriceCash,
			Status:           o.Status,
			Media:            o.Media,
		}
	}

	return offers
}
