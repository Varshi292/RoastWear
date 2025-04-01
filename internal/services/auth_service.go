// Package services ...
package services

import (
	"errors"
	"log"
	"time"

	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repository/interfaces"
	"github.com/Varshi292/RoastWear/internal/session"
	_ "github.com/Varshi292/RoastWear/internal/session"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// AuthService ...
//
// Fields:
//   - repo: ...

// NewAuthService ...
//
// Parameters:
//   - repo: ...
//
// Returns:
//   - *AuthService: ...
type AuthService struct {
	repo           interfaces.UserRepository
	sessionService *SessionService
}

func NewAuthService(repo interfaces.UserRepository, sessionService *SessionService) *AuthService {
	return &AuthService{
		repo:           repo,
		sessionService: sessionService,
	}
}

// LoginUser ...
//
// Parameters:
//   - request: ...
//   - c: ...
//
// Returns:
//   - error: ...
func (service *AuthService) LoginUser(request *models.UserLoginRequest, c *fiber.Ctx) error {
	user, err := service.repo.GetUser("username", request.Username)
	if err != nil || !utils.VerifyPassword(request.Password, user.Password) {
		return utils.ErrInvalidCredentials
	}

	// Always create a new session
	sess, err := session.Store.Get(c)
	if err != nil {
		return utils.ErrSessionNotFound
	}

	// Optionally destroy the old session first (clears existing cookie)
	_ = sess.Destroy()

	// Create a brand new session
	newSess, err := session.Store.Get(c)
	if err != nil {
		return utils.ErrSessionNotFound
	}

	newSess.Set("userID", user.ID)
	newSess.Set("username", user.Username)
	newSess.Set("loginTime", time.Now().Unix())

	if err := newSess.Save(); err != nil {
		return errors.New("failed saving session")
	}

	// ✅ Print login log
	log.Printf("✅ User '%s' logged in with session ID: %s", user.Username, newSess.ID())

	return nil
}
