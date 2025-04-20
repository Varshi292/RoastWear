// Package services ...
package repositories

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"gorm.io/gorm"
)

// SessionRepository handles session operations.
type SessionRepository struct {
	db *gorm.DB
}

// NewSessionRepository creates a new SessionRepository instance.
func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// CreateSession saves a new session to the database.
func (s *SessionRepository) CreateSession(session *models.Session) error {
	return s.db.Create(session).Error
}

// GetSession checks if a session with the given username and session ID exists.
func (s *SessionRepository) GetSession(sessionID string) error {
	sess := &models.Session{}
	if err := s.db.Where("session_key = ?", sessionID).First(sess).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSession removes a session by username and session ID.
func (s *SessionRepository) DeleteSession(username, sessionID string) error {
	return s.db.Where("username = ? AND session_id = ?", username, sessionID).Delete(&models.Session{}).Error
}
