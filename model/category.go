package model

import "github.com/jinzhu/gorm"

// 商品分类
type Category struct {
	gorm.Model
	CategoryName string
}
