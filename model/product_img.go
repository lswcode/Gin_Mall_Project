package model

import "github.com/jinzhu/gorm"

// 商品图片
type ProductImg struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	ImgPath   string
}
