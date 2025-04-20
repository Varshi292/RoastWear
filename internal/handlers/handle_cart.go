package handlers

import (
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CheckoutHandler struct {
	sessionRepository *repositories.SessionRepository
	db                *gorm.DB
}

func NewCheckoutHandler(s *repositories.SessionRepository, db *gorm.DB) *CheckoutHandler {
	return &CheckoutHandler{sessionRepository: s, db: db}
}

func (h *CheckoutHandler) CheckoutCart(c *fiber.Ctx) error {
	var req struct {
		Username  string            `json:"username"`
		SessionID string            `json:"session_id"`
		Items     []models.CartItem `json:"items"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}

	valid, err := h.sessionRepository.VerifySession(req.Username, req.SessionID)
	if err != nil || !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid session",
		})
	}

	for _, item := range req.Items {
		purchase := models.Purchase{
			Username:  req.Username,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		if err := h.db.Create(&purchase).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to store purchase",
			})
		}
	}

	return c.JSON(fiber.Map{"message": "checkout successful"})
}
