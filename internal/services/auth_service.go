// Package services ...
package services

import (
	"github.com/Varshi292/RoastWear/internal/interfaces"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/utils"
)

type AuthService struct {
	userRepo    interfaces.UserRepository
	sessionRepo *repositories.SessionRepository
}

func NewAuthService(repo interfaces.UserRepository, sessionRepo *repositories.SessionRepository) *AuthService {
	return &AuthService{
		userRepo:    repo,
		sessionRepo: sessionRepo,
	}
}

func (service *AuthService) AuthenticateUser(request *models.UserLoginRequest) error {
	user, err := service.userRepo.GetUser("username", request.Username)
	if err != nil || utils.VerifyPassword(request.Password, user.Password) == false {
		return utils.ErrInvalidCredentials
	}
	return nil
}
