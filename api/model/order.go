package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	User     User `gorm:"foreignkey:UserID"`
	UserID   uint
	Quantity uint `json:"quantity"`
}

func (Order) TableName() string {
	return "orders"
}
