package api

import (
	"errors"
	"teklif/internal/api/models"
	response "teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ExpirePropertyOffers(c *gin.Context) {
	var req models.PropertyIDsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	if len(req.PropertyIDs) == 0 {
		response.WriteErrorResponse(c, h.logger, errors.New("property_ids cannot be empty"))
		return
	}

	if err := h.offer.ExpireOffersByProperties(c, req.PropertyIDs); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, nil)
}
