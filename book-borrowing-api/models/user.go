package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phoneNumber"`
}
