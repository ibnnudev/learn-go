package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn" gorm:"unique"`
	Stock  int    `json:"stock"`
}
