package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

var (
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	digitRegex     = regexp.MustCompile(`[0-9]`)
	specialRegex   = regexp.MustCompile(`[^a-zA-Z0-9]`)
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

	_ = cv.RegisterValidation("password", isValidPassword)

	return cv
}

func FieldErrorToString(fe validator.FieldError) string {
	messages := map[string]string{
		"required": "поле обязательно для заполнения",
		"alpha":    "поле должно содержать только буквенные символы",
		"numeric":  "поле должно содержать только цифровые символы",
	}

	if msg, exists := messages[fe.Tag()]; exists {
		return msg
	}

	if fe.Tag() == "min" || fe.Tag() == "max" || fe.Tag() == "gte" || fe.Tag() == "lte" || fe.Tag() == "oneOf" {
		return fmt.Sprintf("поле должно быть %s %s", fe.Tag(), fe.Param())
	}

	return "поле недействительно"
}

func isValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return utf8.RuneCountInString(password) >= 8 &&
		uppercaseRegex.MatchString(password) &&
		digitRegex.MatchString(password) &&
		specialRegex.MatchString(password)
}
