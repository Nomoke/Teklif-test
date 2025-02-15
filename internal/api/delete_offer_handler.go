package api

import (
	response "teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

// @Summary Delete an offer
// @Description Delete an offer by its ID
// @Tags offers
// @Accept  json
// @Produce  json
// @Param   offer_id  path    string  true  "ID of the offer to be deleted"
// @Router /offers/{offerID}/delete [delete]
func (h *handler) DeleteOffer(c *gin.Context) {
	offerID, err := OfferID(c)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	if err := h.offer.DeleteOffer(c, offerID); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, nil)
}
