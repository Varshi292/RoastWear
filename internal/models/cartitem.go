// internal/models/cart_item.go
package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Username  string `gorm:"index:idx_user_product,unique;not null"`
	ProductID string `gorm:"index:idx_user_product,unique;not null"`
	Quantity  int    `gorm:"not null"`
}
