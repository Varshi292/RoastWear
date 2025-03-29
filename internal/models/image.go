package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UploadedBy string `gorm:"not null"`
	Filename   string `gorm:"not null"`
	Data       []byte `gorm:"type:blob;not null"`
	Size       int    `gorm:"not null"`
	Type       string `gorm:"not null"`
}
