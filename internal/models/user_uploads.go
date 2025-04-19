package models

import (
	"gorm.io/gorm"
)

// UserUpload represents lightweight image metadata associated with a user.
// Used when storing images to disk (not as blobs).
type UserUpload struct {
	gorm.Model
	Username string `gorm:"not null"`
	Filepath string `gorm:"not null"`
}
