package api

import (
	"errors"
	"teklif/internal/api/models"
	response "teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetActiveOffersByProperties(c *gin.Context) {
	var req models.PropertyIDsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	if len(req.PropertyIDs) == 0 {
		response.WriteErrorResponse(c, h.logger, errors.New("property_ids cannot be empty"))
		return
	}

	offers, err := h.offer.GetActiveOffersByProperties(c, req.PropertyIDs)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, gin.H{
		"offers": models.BuildOffersResponse(offers),
	})
}
