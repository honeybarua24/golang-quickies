package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" validate:"required,min=3,max=32"`
	Email  string `json:"email" validate:"email"`
	Number uint   `json:"number"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
