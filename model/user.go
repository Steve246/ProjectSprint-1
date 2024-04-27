package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"unique" json:"userName"`
	Email            string `json:"userEmail"`
	Password         string `json:"userPassword"`
	RegistrationDate string `json:"registration_date"`
}
