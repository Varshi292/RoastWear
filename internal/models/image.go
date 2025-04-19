// Package models ...
package models

import (
	"gorm.io/gorm"
)

// Image ...
//
// Fields:
//   - ID: ...
//   - CreatedAt: ...
//   - UpdatedAt: ...
//   - DeletedAt: ...
//   - UploadedBy: ...
//   - Filename: ...
//   - Data: ...
//   - Size: ...
//   - Type: ...
type Image struct {
	gorm.Model
	UploadedBy string `gorm:"not null"`
	Filename   string `gorm:"not null"`
	Data       []byte `gorm:"type:blob;not null"`
	Size       int    `gorm:"not null"`
	Type       string `gorm:"not null"`
}
