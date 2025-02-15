package validator

import (
	"fmt"
	"reflect"
	"strings"
	"teklif/internal/api/models"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	*validator.Validate
}

func New() *CustomValidator {
	cv := &CustomValidator{validator.New()}

	cv.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	cv.RegisterStructValidation(models.OfferValidator, models.CreateOfferReq{})

	return cv
}

func FieldErrorToString(fe validator.FieldError) string {
	messages := map[string]string{
		"required":                      "поле обязательно для заполнения",
		"alpha":                         "поле должно содержать только буквенные символы",
		"numeric":                       "поле должно содержать только цифровые символы",
		"required_for_extra_commission": "поле обязательно для комиссии",
		"required_for_discount":         "поле обязательно для скидки",
		"required_for_special_price":    "поле обязательно для спец. цены",
	}

	if msg, exists := messages[fe.Tag()]; exists {
		return msg
	}

	if fe.Tag() == "min" || fe.Tag() == "max" || fe.Tag() == "gte" || fe.Tag() == "lte" || fe.Tag() == "oneOf" {
		return fmt.Sprintf("поле должно быть %s %s", fe.Tag(), fe.Param())
	}

	return "поле недействительно"
}
