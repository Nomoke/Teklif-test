package models

import "github.com/go-playground/validator/v10"

const (
	OfferTypeGift            = "gift"
	OfferTypeExtraCommission = "extra_commission"
	OfferTypeDiscount        = "discount"
	OfferTypeSpecialPrice    = "special_price"
)

func OfferValidator(sl validator.StructLevel) {
	offer := sl.Current().Interface().(CreateOfferReq)

	switch offer.Type {
	case OfferTypeGift:

	case OfferTypeExtraCommission:
		if offer.CommissionAgent == nil || offer.CommissionAgency == nil {
			sl.ReportError(offer.CommissionAgent, "commission_agent", "CommissionAgent", "required_for_extra_commission", "")
			sl.ReportError(offer.CommissionAgency, "commission_agency", "CommissionAgency", "required_for_extra_commission", "")
		}

	case OfferTypeDiscount:
		if offer.Discount == nil {
			sl.ReportError(offer.Discount, "discount", "Discount", "required_for_discount", "")
		}

	case OfferTypeSpecialPrice:
		if offer.PriceCash == nil && offer.PriceInstallment == nil {
			sl.ReportError(offer.PriceCash, "price_cash", "PriceCash", "required_for_special_price", "")
			sl.ReportError(offer.PriceInstallment, "price_installment", "PriceInstallment", "required_for_special_price", "")
		}
	}
}
