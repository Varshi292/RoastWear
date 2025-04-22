// internal/repositories/purchase_repository.go
package repositories

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"gorm.io/gorm"
)

type PurchaseRepository struct {
	DB *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{DB: db}
}

func (r *PurchaseRepository) CreatePurchase(username string, products string, total string) error {
	purchase := models.Purchase{
		Username: username,
		Products: products,
		Total:    total,
	}
	return r.DB.Create(&purchase).Error
}

func (r *PurchaseRepository) GetPurchasesByUsername(username string) ([]models.Purchase, error) {
	var purchases []models.Purchase
	err := r.DB.Where("username = ?", username).Find(&purchases).Error
	return purchases, err
}
