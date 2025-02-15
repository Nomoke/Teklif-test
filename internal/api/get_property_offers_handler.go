package api

import (
	"errors"
	"teklif/internal/api/models"
	"teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get active offers by property IDs
// @Description Retrieve a list of active offers based on property IDs that are approved and not expired
// @Tags offers
// @Accept  json
// @Produce  json
// @Param property_ids body models.PropertyIDsReq true "List of property IDs"
// @Success 200 {object} []models.OfferResp "List of active offers"
// @Failure 400 {object} http.ErrorResponse "Bad request error"
// @Failure 404 {object} http.ErrorResponse "No offers found"
// @Failure 500 {object} http.ErrorResponse "Internal server error"
// @Router /offers/properties [post]
func (h *handler) GetActiveOffersByProperties(c *gin.Context) {
	var req models.PropertyIDsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	if len(req.PropertyIDs) == 0 {
		http.WriteErrorResponse(c, h.logger, errors.New("property_ids cannot be empty"))
		return
	}

	offers, err := h.offer.GetActiveOffersByProperties(c, req.PropertyIDs)
	if err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	http.WriteSuccessResponse(c, gin.H{
		"offers": models.BuildOffersResponse(offers),
	})
}
