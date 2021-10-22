package model

type OrderProduct struct {
	Order     Order `gorm:"foreignkey:OrderID"`
	OrderID   uint
	Product   Product `gorm:"foreignkey:ProductID"`
	ProductID uint
}
