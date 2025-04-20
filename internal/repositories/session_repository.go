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
func (repo *SessionRepository) CreateSession(session *models.Session) error {
	return repo.db.Create(session).Error
}

// GetSession checks if a session with the session ID exists.
func (repo *SessionRepository) GetSession(sessionID string) error {
	sess := &models.Session{}
	if err := repo.db.Where("session_key = ?", sessionID).First(sess).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSession removes a session by session ID.
func (repo *SessionRepository) DeleteSession(session *models.Session) error {
	return repo.db.Delete(session).Error
}
