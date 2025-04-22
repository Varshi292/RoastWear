package models

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

// Session is the actual GORM + Fiber session model used in code
type Session struct {
	gorm.Model
	*session.Session `gorm:"-"`
	SessionKey       string `gorm:"uniqueIndex;not null"`
}

// SessionDoc is a flattened version used only for Swagger documentation
// @Description This is the session model used for Swagger documentation
type SessionDoc struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SessionKey string    `json:"session_key"`
}

// GenericResponse represents a basic API response format
type GenericResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
