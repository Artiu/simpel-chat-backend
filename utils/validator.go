package validator

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func New() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("hasUppercase", hasUppercase)
	v.RegisterValidation("hasLowercase", hasLowercase)
	v.RegisterValidation("hasNumber", hasNumber)
	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func hasUppercase(field validator.FieldLevel) bool {
	text := field.Field().String()
	return strings.ToLower(text) != text
}

func hasLowercase(field validator.FieldLevel) bool {
	text := field.Field().String()
	return strings.ToUpper(text) != text
}

func hasNumber(field validator.FieldLevel) bool {
	text := field.Field().String()
	for _, value := range strings.Split(text, "") {
		_, err := strconv.Atoi(value)
		if err == nil {
			return true
		}
	}
	return false
}
