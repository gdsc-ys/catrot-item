package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Title    string
	Category string
	Price    int
	Content  string
}
