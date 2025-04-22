package repositories

import (
	"errors"
	"fmt"

	"github.com/Varshi292/RoastWear/internal/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

// ModifyItem updates or inserts a cart item, calculating total price from unit price * quantity.
func (r *CartRepository) ModifyItem(username string, productID int, quantity int, totalPrice float64) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}

	if quantity == 0 {
		return r.DB.
			Unscoped().
			Where("username = ? AND product_id = ?", username, productID).
			Delete(&models.CartItem{}).Error
	}

	var item models.CartItem
	err := r.DB.
		Where("username = ? AND product_id = ?", username, productID).
		First(&item).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newItem := models.CartItem{
				Username:   username,
				ProductID:  productID,
				Quantity:   quantity,
				TotalPrice: fmt.Sprintf("%.2f", totalPrice),
			}
			return r.DB.Create(&newItem).Error
		}
		return err
	}

	item.Quantity = quantity
	item.TotalPrice = fmt.Sprintf("%.2f", totalPrice)
	return r.DB.Save(&item).Error
}

func (r *CartRepository) DeleteItem(username string, productID int) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	return r.DB.
		Unscoped().
		Where("username = ? AND product_id = ?", username, productID).
		Delete(&models.CartItem{}).Error
}

func (r *CartRepository) GetItemsByUsername(username string) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.DB.Where("username = ?", username).Find(&items).Error
	return items, err
}

func (r *CartRepository) ClearCartForUser(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	return r.DB.
		Unscoped().
		Where("username = ?", username).
		Delete(&models.CartItem{}).Error
}
