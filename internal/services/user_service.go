// Package services ...
package services

import (
	"RoastWear/internal/interfaces"
	"RoastWear/internal/models"
	"RoastWear/internal/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

// UserService ...
//
// Fields:
//   - repo: ...
type UserService struct {
	repo interfaces.UserRepository
}

// NewUserService ...
//
// Parameters:
//   - repo: ...
//
// Returns:
//   - *UserService: ...
func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// RegisterUser ...
//
// Parameters:
//   - request: ...
//
// Returns:
//   - error: ...
func (service *UserService) RegisterUser(request *models.UserRegisterRequest) error {
	if service.repo.HasUser("username", request.Username) {
		return utils.NewErrUserExists(request.Username)
	}
	if service.repo.HasUser("email", request.Email) {
		return utils.NewErrEmailExists(request.Email)
	}
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}
	user := &models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}
	if err := service.repo.CreateUser(user); err != nil {
		return fmt.Errorf("error creating user '%s': %v", request.Username, err)
	}
	log.Printf("✅ User '%s' created successfully.\n", user.Username)
	return nil
}

// RemoveUser ...
//
// Parameters:
//   - username: ...
//
// Returns:
//   - error: ...
func (service *UserService) RemoveUser(username string) error {
	target, err := service.repo.GetUser("username", username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("user '%s' not found", username)
		}
		return fmt.Errorf("error retrieving user '%s': %v", username, err)
	}
	if err := service.repo.DeleteUser(target); err != nil {
		return fmt.Errorf("error removing user '%s': %v", username, err)
	}
	log.Printf("✅ User '%s' deleted successfully.\n", username)
	return nil
}

// ShowUsers ...
//
// Parameters:
//   - none
//
// Returns:
//   - error: ...
func (service *UserService) ShowUsers() error {
	users, err := service.repo.GetAllUsers()
	if err != nil {
		return err
	}
	if len(users) == 0 {
		log.Println("No users found in the database.")
		return nil
	}
	log.Println("User Data:")
	log.SetFlags(0)
	log.Printf("%-5s | %-20s | %-30s | %-20s | %-30s\n", "ID", "Username", "Email", "Created At", "Updated At")
	log.Println("---------------------------------------------------------------------------------------------------------")
	for _, user := range users {
		log.Printf("%-5d | %-20s | %-30s | %-20s | %-30s\n",
			user.ID, user.Username, user.Email,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
			user.UpdatedAt.Format("2006-01-02 15:04:05"),
		)
	}
	return nil
}
