package models

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	*session.Session
}
