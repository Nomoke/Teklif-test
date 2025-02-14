package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OfferID(c *gin.Context) (uuid.UUID, error) {
	oIDStr := c.Param("offerID")

	if oIDStr == "" {
		return uuid.UUID{}, errors.New("offer ID is required")
	}

	id, err := uuid.Parse(oIDStr)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid offer ID format")
	}

	return id, nil
}
