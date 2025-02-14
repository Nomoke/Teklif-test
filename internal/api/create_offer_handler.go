package api

import (
	"github.com/gin-gonic/gin"

	response "teklif/internal/pkg/http"

	"teklif/internal/api/models"
	"teklif/internal/models/domain"
	"teklif/internal/pkg/storage"
)

func (h *handler) CreateOffer(c *gin.Context) {
	var req models.CreateOfferReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	filePaths, err := storage.SaveFiles(c, form.File["media"])
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	offer := domain.Offer{
		Title:            req.Title,
		Description:      req.Description,
		Type:             req.Type,
		ExpiryDate:       req.ExpiryDate,
		CommissionAgent:  req.CommissionAgent,
		CommissionAgency: req.CommissionAgency,
		Discount:         req.Discount,
		PriceInstallment: req.PriceInstallment,
		PriceCash:        req.PriceCash,
		Status:           "pending",
		Media:            filePaths,
	}

	if err := h.offer.CreateOffer(c, offer); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	response.WriteSuccessResponse(c, nil)
}
