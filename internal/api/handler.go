package api

import (
	"github.com/gin-gonic/gin"

	validatorpkg "teklif/internal/pkg/validator"
)

type handler struct {
	logger    logger
	offer     offerService
	validator *validatorpkg.CustomValidator
}

func New(
	logger logger,
	offer offerService,
) (*gin.Engine, error) {
	v := validatorpkg.New()

	h := &handler{
		logger:    logger,
		offer:     offer,
		validator: v,
	}

	r := gin.New()

	r.Use(gin.Recovery())

	api := r.Group("/api")

	api.POST("/offers", h.CreateOffer)
	api.GET("/offers/active", h.GetActiveOffers)
	api.GET("/offers/week-expired", h.GetWeekExpiredOffers)
	api.GET("/offers/properties", h.GetActiveOffersByProperties)
	api.POST("/offers/attach-properties", h.AttachPropertyOffer)
	api.POST("/offers/delete", h.DeleteOffer)
	api.POST("/offers/approve", h.ApproveOffer)
	api.POST("/offers/expire-properties", h.ExpirePropertyOffers)

	return r, nil
}
