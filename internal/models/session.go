package models

import (
	"gorm.io/gorm"
)

// Session ...
//
// Fields:
//   - ID: Auto-increment primary key
//   - CreatedAt, UpdatedAt, DeletedAt: From gorm.Model
//   - Username: The user's unique name
//   - SessionID: The session identifier (e.g., UUID or token)
type Session struct {
	gorm.Model
	Username  string `gorm:"not null;index"`
	SessionID string `gorm:"not null;uniqueIndex"`
}
