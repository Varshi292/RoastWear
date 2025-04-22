package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Username   string `gorm:"index:idx_user_product,unique;not null"`
	ProductID  int    `gorm:"index:idx_user_product,unique;not null"`
	Quantity   int    `gorm:"not null"`
	TotalPrice string // Store the total price for this row (quantity * unit price)
}
