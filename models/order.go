package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	Status     string      `json:"status"` // "pending", "processed", "shipped", "delivered"
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `json:"product" gorm:"foreignkey:ProductID"`
}
