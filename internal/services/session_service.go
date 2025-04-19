// Package services ...
package services

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"gorm.io/gorm"
)

// SessionService handles session operations.
type SessionService struct {
	db *gorm.DB
}

// NewSessionService creates a new SessionService instance.
func NewSessionService(db *gorm.DB) *SessionService {
	return &SessionService{db: db}
}

// CreateSession saves a new session to the database.
func (s *SessionService) CreateSession(session *models.Session) error {
	return s.db.Create(session).Error
}

// VerifySession checks if a session with the given username and session ID exists.
func (s *SessionService) VerifySession(username, sessionID string) (bool, error) {
	var session models.Session
	err := s.db.Where("username = ? AND session_id = ?", username, sessionID).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// DeleteSession removes a session by username and session ID.
func (s *SessionService) DeleteSession(username, sessionID string) error {
	return s.db.Where("username = ? AND session_id = ?", username, sessionID).Delete(&models.Session{}).Error
}
