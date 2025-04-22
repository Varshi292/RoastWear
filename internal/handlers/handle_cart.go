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

// CheckoutCart handles the checkout process for a user's shopping cart.
// @Summary Checkout cart
// @Description This endpoint allows a user to complete the checkout process by providing a session ID, username, and the items in their cart.
// @Tags Cart
// @Accept json
// @Produce json
// @Param checkoutRequest body struct { Username string `json:"username"`; SessionID string `json:"session_id"`; Items []models.CartItem `json:"items"` } true "Checkout request details"
// @Success 200 {object} map[string]string {"message": "checkout successful"}
// @Failure 400 {object} map[string]string {"error": "invalid request format"}
// @Failure 401 {object} map[string]string {"error": "invalid session"}
// @Failure 500 {object} map[string]string {"error": "failed to store purchase"}
// @Router /checkout [post]
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

	err := h.sessionRepository.GetSession(req.SessionID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid session",
		})
	}

	for _, item := range req.Items {
		purchase := models.Purchase{
			Username:  "",
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
