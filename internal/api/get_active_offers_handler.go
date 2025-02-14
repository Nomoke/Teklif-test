package api

import (
	response "teklif/internal/pkg/http"

	"teklif/internal/api/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetActiveOffers(c *gin.Context) {
	offers, err := h.offer.GetActiveOffers(c)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, gin.H{
		"offers": models.BuildOffersResponse(offers),
	})
}
