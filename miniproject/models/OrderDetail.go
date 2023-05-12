package models

import (
	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id" form:"user_id"`
	ProductID   uint   `json:"product_id" form:"product_id"`
	Quantity    uint   `json:"quantity" form:"quantity"`
	OrderStatus string `json:"status"`
	TotalPrice  uint   `json:"total_price"`
}
