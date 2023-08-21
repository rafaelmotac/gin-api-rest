package util

import (
	"api-go-gin/config"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ValidateModel(s interface{}) []string {
	err := config.Validate.Struct(s)
	if err != nil {
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("Field '%s': %s=%s", err.Field(), err.Tag(), err.Param()))
		}

		return errs
	}

	return nil

}
