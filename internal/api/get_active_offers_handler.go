package api

import (
	"teklif/internal/pkg/http"

	"teklif/internal/api/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get active offers
// @Description Retrieve a list of active offers that are approved and not expired
// @Tags offers
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.OfferResp "List of active offers"
// @Failure 400 {object} http.ErrorResponse "Bad request error"
// @Failure 404 {object} http.ErrorResponse "No offers found"
// @Failure 500 {object} http.ErrorResponse "Internal server error"
// @Router /offers/active [get]
func (h *handler) GetActiveOffers(c *gin.Context) {
	offers, err := h.offer.GetActiveOffers(c)
	if err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	http.WriteSuccessResponse(c, gin.H{
		"offers": models.BuildOffersResponse(offers),
	})
}
