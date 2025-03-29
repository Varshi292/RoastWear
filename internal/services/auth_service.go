package services

import (
	"errors"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repository/interfaces"
	"github.com/Varshi292/RoastWear/internal/session"
	_ "github.com/Varshi292/RoastWear/internal/session"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthService struct {
	repo interfaces.UserRepository
}

func NewAuthService(repo interfaces.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (service *AuthService) LoginUser(request *models.UserLoginRequest, c *fiber.Ctx) error {
	user, err := service.repo.GetUser("username", request.Username)
	if err != nil || !utils.VerifyPassword(user.Password, request.Password) {
		return errors.New("invalid username or password")
	}
	sess, err := session.Store.Get(c)
	if err != nil {
		return errors.New("session not found")
	}
	sess.Set("userID", user.ID)
	sess.Set("username", user.Username)
	sess.Set("loginTime", time.Now().Unix())
	if err := sess.Save(); err != nil {
		return errors.New("failed saving session")
	}
	return nil
}
