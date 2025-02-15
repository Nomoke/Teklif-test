package api

import (
	"errors"
	"teklif/internal/api/models"
	"teklif/internal/models/domain"
	response "teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

// @Summary Attach properties to an offer
// @Description Attach properties to a specific offer by their IDs
// @Tags offers
// @Accept json
// @Produce json
// @Param offer body models.AttachOfferPropertiesReq true "Request body for attaching properties"
// @Example {"offer_id": "123e4567-e89b-12d3-a456-426614174000", "property_ids": ["123e4567-e89b-12d3-a456-426614174001", "123e4567-e89b-12d3-a456-426614174002"]}
// @Success 200 {object} string "Properties successfully attached"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Internal server error"
// @Router /offers/attach-properties [post]
func (h *handler) AttachPropertyOffer(c *gin.Context) {
	var req models.AttachOfferPropertiesReq

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
