package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	Username  string `gorm:"not null"`
	ProductID string `gorm:"not null"`
	Quantity  int    `gorm:"not null"`
}
