package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"log"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	err := Validate.RegisterValidation("not-blank", validators.NotBlank)
	if err != nil {
		log.Fatalf("Error registering validation: %v", err)
	}
}
