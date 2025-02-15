package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"teklif/internal/pkg/http"
	"teklif/internal/pkg/storage"

	"teklif/internal/api/models"
	"teklif/internal/models/domain"
)

// @Summary Create a new offer
// @Description Create a new offer with media files (uploaded as form-data)
// @Tags offers
// @Accept  multipart/form-data
// @Produce  json
// @Param title formData string true "Offer Title"
// @Param description formData string true "Offer Description"
// @Param type formData string true "Offer Type"
// @Param expiry_date formData string true "Offer Expiry Date" format(date)
// @Param commission_agent formData number false "Commission Agent"
// @Param commission_agency formData number false "Commission Agency"
// @Param discount formData number false "Discount"
// @Param price_installment formData number false "Price Installment"
// @Param price_cash formData number false "Price Cash"
// @Param media formData file true "Media files"
// @Success 200 {object} string "Offer created successfully"
// @Failure 400 {object} http.ErrorResponse "Invalid input"
// @Failure 500 {object} http.ErrorResponse "Internal server error"
// @Router /offers [post]
func (h *handler) CreateOffer(c *gin.Context) {
	var req models.CreateOfferReq

	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	var filePaths []string
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["media"]
		filePaths, err = storage.SaveFiles(c, files)
		if err != nil {
			http.WriteErrorResponse(c, h.logger, err)
			return
		}
	}

	if err := h.validator.Struct(req); err != nil {
		http.WriteErrorResponse(c, h.logger, err)
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
		Status:           models.OfferStatusPending,
		Media:            filePaths,
	}

	if err := h.offer.CreateOffer(c, offer); err != nil {
		http.WriteErrorResponse(c, h.logger, err)
		return
	}

	http.WriteSuccessResponse(c, nil)
}
