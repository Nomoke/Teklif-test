package api

import (
	"teklif/internal/api/models"
	"teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get list of offers with "Expired" status that have expired in the last 7 days.
// @Description Retrieve a list of offers that are marked as "Expired" and expired in the last 7 days.
// @Tags offers
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.OfferResp "List of expired offers"
// @Failure 400 {object} http.ErrorResponse "Bad request error"
// @Failure 404 {object} http.ErrorResponse "No offers found"
// @Failure 500 {object} http.ErrorResponse "Internal server error"
// @Router /offers/week-expired [get]
func (h *handler) GetWeekExpiredOffers(c *gin.Context) {
	offers, err := h.offer.GetWeekExpiredOffers(c)
	if err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	http.WriteSuccessResponse(c, gin.H{
		"offers": models.BuildOffersResponse(offers),
	})
}
