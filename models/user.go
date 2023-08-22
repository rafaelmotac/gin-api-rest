package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `gorm:"size:255;not-null;unique"`
	Password   string `gorm:"size:255;not-null"`
}

type UserDTO struct {
	Username string `json:"username,omitempty" validate:"required=true,not-blank=true,max=255"`
	Password string `json:"password,omitempty" validate:"required=true,not-blank=true,max=255"`
}

func (user *User) ToDTO() *UserDTO {
	return &UserDTO{
		Username: user.Username,
		Password: user.Password,
	}
}

func (user *UserDTO) ToModel() *User {
	return &User{
		Username: user.Username,
		Password: user.Password,
	}
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = user.HashPassword(); err != nil {
		return err
	}
	return nil
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	if user.Password != "" {
		user.Password = "*****"
	}
	return nil
}

func (user *User) AfterFind(tx *gorm.DB) (err error) {
	if user.Password != "" {
		user.Password = "*****"
	}
	return
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return nil
}
