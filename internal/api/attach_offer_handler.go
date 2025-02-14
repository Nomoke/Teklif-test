package api

import (
	"errors"
	api "teklif/internal/api/models"
	"teklif/internal/models/domain"
	response "teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AttachPropertyOffer(c *gin.Context) {
	var req api.AttachOfferPropertiesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	if len(req.PropertyIDs) == 0 {
		response.WriteErrorResponse(c, h.logger, errors.New("property_ids cannot be empty"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	attach := domain.AttachOffer{
		OfferID:     req.OfferID,
		PropertyIDs: req.PropertyIDs,
	}

	if err := h.offer.AttachOffer(c, attach); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, nil)
}
