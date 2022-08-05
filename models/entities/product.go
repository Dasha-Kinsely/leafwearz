package entities

import "gorm.io/gorm"

type ProductBase struct {
	gorm.Model
	Title string `json:"title" gorm:"column:title;unique"`
	CoverImage Image
	ImageList []Image
	VideoList []Video
	BasePrice float32 `json:"base_price" gorm:"column:base_price"`
	Manufacturer string `json:"manufacturer" gorm:"column:manufacturer"`
	Designer string `json:"designer" gorm:"column:designer"`
}



