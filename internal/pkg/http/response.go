package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	validatorpkg "teklif/internal/pkg/validator"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error  any    `json:"error"`
}

type ValidationErrors struct {
	Status string       `json:"status"`
	Errors []FieldError `json:"errors"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type logger interface {
	Error(text ...any)
}

func WriteSuccessResponse(c *gin.Context, data any) {
	if data == nil {
		c.Status(http.StatusOK)
		return
	}
	c.JSON(http.StatusOK, data)
}

func WriteErrorResponse(c *gin.Context, logger logger, err error) {
	if vErr, ok := errors.Cause(err).(validator.ValidationErrors); ok {
		errs := make([]FieldError, len(vErr))
		for i, fe := range vErr {
			errs[i] = FieldError{
				Field:   fe.Field(),
				Message: validatorpkg.FieldErrorToString(fe),
			}
		}
		c.JSON(http.StatusUnprocessableEntity, ValidationErrors{"errors", errs})
		return
	}

	logger.Error(err)
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Status: "error",
		Error:  "Что-то пошло не так, попробуйте еще раз",
	})
}
