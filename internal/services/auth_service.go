// Package services ...
package services

import (
	"RoastWear/internal/interfaces"
	"RoastWear/internal/models"
	"RoastWear/internal/session"
	"RoastWear/internal/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"time"
)

// AuthService ...
//
// Fields:
//   - repo: ...
type AuthService struct {
	repo interfaces.UserRepository
}

// NewAuthService ...
//
// Parameters:
//   - repo: ...
//
// Returns:
//   - *AuthService: ...
func NewAuthService(repo interfaces.UserRepository) *AuthService {
	return &AuthService{repo: repo}
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
	sess, err := session.Store.Get(c)
	if err != nil {
		return utils.ErrSessionNotFound
	}
	sess.Set("userID", user.ID)
	sess.Set("username", user.Username)
	sess.Set("loginTime", time.Now().Unix())
	if err := sess.Save(); err != nil {
		return errors.New("failed saving session")
	}
	return nil
}
