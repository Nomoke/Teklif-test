package api

import (
	response "teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ApproveOffer(c *gin.Context) {
	offerID, err := OfferID(c)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	if err := h.offer.ApproveOffer(c, offerID); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, nil)
}
