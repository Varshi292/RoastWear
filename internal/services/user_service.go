package services

import (
	"errors"
	"fmt"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repository/interfaces"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(request *models.UserCreateRequest) error {
	if s.repo.HasUser("username", request.Username) {
		return fmt.Errorf("username '%s' already exists", request.Username)
	}
	if s.repo.HasUser("email", request.Email) {
		return fmt.Errorf("email '%s' already exists", request.Email)
	}
	user := &models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return fmt.Errorf("error creating user '%s': %v", request.Username, err)
	}
	log.Printf("✅ User '%s' created successfully.\n", user.Username)
	return nil
}

func (s *UserService) RemoveUser(username string) error {
	target, err := s.repo.GetUser("username", username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("user '%s' not found", username)
		}
		return fmt.Errorf("error retrieving user '%s': %v", username, err)
	}
	if err := s.repo.DeleteUser(target); err != nil {
		return fmt.Errorf("error removing user '%s': %v", username, err)
	}
	log.Printf("✅ User '%s' deleted successfully.\n", username)
	return nil
}
