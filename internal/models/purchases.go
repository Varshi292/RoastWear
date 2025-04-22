package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	Username string `gorm:"not null"`
	Products string `gorm:"type:text;not null"` // Format: productID:quantity#productID:quantity
	Total    string `gorm:"not null"`           // Total purchase amount as formatted string (e.g., "49.99")
}
