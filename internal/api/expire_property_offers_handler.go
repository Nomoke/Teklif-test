package api

import (
	"errors"
	"teklif/internal/api/models"
	"teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

// @Summary Expire offers by property IDs
// @Description Expire all offers associated with the provided property IDs
// @Tags offers
// @Accept  json
// @Produce  json
// @Param property_ids body models.PropertyIDsReq true "Request body containing property IDs"
// @Router /offers/expire-properties [post]
func (h *handler) ExpirePropertyOffers(c *gin.Context) {
	var req models.PropertyIDsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	if len(req.PropertyIDs) == 0 {
		http.WriteErrorResponse(c, h.logger, errors.New("property_ids cannot be empty"))
		return
	}

	if err := h.offer.ExpireOffersByProperties(c, req.PropertyIDs); err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	http.WriteSuccessResponse(c, nil)
}
