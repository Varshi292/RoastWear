package interfaces

import "github.com/Varshi292/RoastWear/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	DeleteUser(user *models.User) error
	HasUser(field string, value interface{}) bool
	GetUser(field string, value interface{}) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}
