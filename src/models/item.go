package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Title    string `json:"title" form:"title"`
	Category string `json:"category" form:"catetory"`
	Price    int `json:"price" form:"price"`
	Content  string `json:"content" form:"content"`
}
