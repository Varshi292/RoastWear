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

// UploadSuccessResponse represents the response after a successful image upload
type UploadSuccessResponse struct {
	Message  string `json:"message"`
	Filepath string `json:"filepath"`
}

// ErrorResponse represents an error response for any failed request
type ErrorResponse struct {
	Error string `json:"error"`
}
