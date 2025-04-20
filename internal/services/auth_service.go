// Package services ...
package services

import (
	"github.com/Varshi292/RoastWear/internal/interfaces"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

type AuthService struct {
	userRepo    interfaces.UserRepository
	sessionRepo *repositories.SessionRepository
}

func NewAuthService(repo interfaces.UserRepository, sessionService *repositories.SessionRepository) *AuthService {
	return &AuthService{
		userRepo:    repo,
		sessionRepo: sessionService,
	}
}

func (service *AuthService) LoginUser(request *models.UserLoginRequest, c *fiber.Ctx) (*session.Session, error) {
	user, err := service.userRepo.GetUser("username", request.Username)
	if err != nil || !utils.VerifyPassword(request.Password, user.Password) {
		return nil, utils.ErrInvalidCredentials
	}

	// Always create a new session
	sess, err := sessions.Store.Get(c)
	if err != nil {
		return nil, utils.ErrSessionNotFound
	}

	// Optionally destroy the old session first (clears existing cookie)
	_ = sess.Destroy()
	// Create a new session
	newSess, err := sessions.Store.Get(c)
	if err != nil {
		return nil, utils.ErrSessionNotFound
	}
	newSess.Set("userID", user.ID)
	newSess.Set("username", user.Username)
	newSess.Set("loginTime", time.Now().Unix())
	return newSess, nil
}
