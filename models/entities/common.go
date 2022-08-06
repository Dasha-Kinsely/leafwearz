package entities

import "gorm.io/gorm"

// Note: size column is always in KB, it will be rounded up by util functions is needed

type Image struct {
	gorm.Model
	Label string `json:"label" gorm:"column:label;notNull"` 
	Size int64 `json:"size" gorm:"column:img_size"`
	Address string `gorm:"column:address;notNull"`
}

type Video struct {
	gorm.Model
	Label string `json:"label" gorm:"column:label;notNull"` 
	Size int64 `json:"size" gorm:"column:img_size"` 
	Address string `gorm:"column:address;notNull"`
}

type Empty struct {}