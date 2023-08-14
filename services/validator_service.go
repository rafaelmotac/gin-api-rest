package services

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterValidation("not-blank", validators.NotBlank)
}

func ValidateModel(s interface{}) []string {
	err := Validate.Struct(s)
	if err != nil {
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("Field '%s': %s=%s", err.Field(), err.Tag(), err.Param()))
		}

		return errs
	}

	return nil

}
