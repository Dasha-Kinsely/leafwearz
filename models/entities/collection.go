package entities

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Title string `json:"title" gorm:"column:title"`
	AdminLabel string `json:"admin_label" gorm:"column:admin_label"`
	Description string `json:"description" gorm:"column:description"`
	ProductList []ProductBase
}