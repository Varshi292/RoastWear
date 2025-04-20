// Package repository ...
package repositories

import (
	"fmt"
	"github.com/Varshi292/RoastWear/internal/models"
	"gorm.io/gorm"
)

// UserRepository ...
//
// Fields:
//   - Db: ...
type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.Db.Create(&user).Error
}

func (repo *UserRepository) DeleteUser(user *models.User) error {
	return repo.Db.Delete(user).Error
}

func (repo *UserRepository) HasUser(field string, value interface{}) bool {
	_, err := repo.GetUser(field, value)
	return err == nil
}

func (repo *UserRepository) GetUser(field string, value interface{}) (*models.User, error) {
	target := &models.User{}
	err := repo.Db.Where(fmt.Sprintf("%s = ?", field), value).First(&target).Error
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
