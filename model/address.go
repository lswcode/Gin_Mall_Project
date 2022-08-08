package model

import "github.com/jinzhu/gorm"

// 用户收货地址
type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"type:varchar(50) not null"`
}
