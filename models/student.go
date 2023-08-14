package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model           `swaggerignore:"true"`
	Name                 string `json:"name,omitempty" validate:"required=true,not-blank=true"`
	FiscalNumber         string `json:"fiscalNumber,omitempty" validate:"number=true,required=true,not-blank=true,min=11"`
	IdentificationNumber string `json:"identificationNumber,omitempty" validate:"number=true,required=true,not-blank=true,min=9"`
}
