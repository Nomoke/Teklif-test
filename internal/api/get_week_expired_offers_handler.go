package api

import (
	"teklif/internal/api/models"
	response "teklif/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetWeekExpiredOffers(c *gin.Context) {
	offers, err := h.offer.GetWeekExpiredOffers(c)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, gin.H{
		"offers": models.BuildOffersResponse(offers),
	})
}
