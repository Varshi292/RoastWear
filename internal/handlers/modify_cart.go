package handlers

import (
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

// ModifyCartRequest represents the request body to modify a cart item
type ModifyCartRequest struct {
	Username  string  `json:"username"`
	ProductID int     `json:"productid"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unitPrice"`
}

// MessageResponse is a generic response with a message
type MessageResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success,omitempty"`
}

// CartHandler handles cart operations
type CartHandler struct {
	Repo *repositories.CartRepository
}

// NewCartHandler initializes a new CartHandler
func NewCartHandler(repo *repositories.CartRepository) *CartHandler {
	return &CartHandler{Repo: repo}
}

// ModifyCart modifies the user's cart (add, update, or delete items).
// @Summary      Modify user's cart
// @Description  Adds an item to the user's cart, updates quantity, or deletes it if quantity is 0.
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        cartItem body handlers.ModifyCartRequest true "Cart item details"
// @Success      200 {object} handlers.MessageResponse
// @Failure      400 {object} handlers.MessageResponse
// @Failure      500 {object} handlers.MessageResponse
// @Router       /cart/modify [post]
func (h *CartHandler) ModifyCart(c *fiber.Ctx) error {
	var body ModifyCartRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(MessageResponse{
			Message: "Invalid request",
		})
	}

	if body.Username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(MessageResponse{
			Message: "You must be logged in to store cart information.",
		})
	}

	// Delete if quantity is 0
	if body.Quantity == 0 {
		err := h.Repo.DeleteItem(body.Username, body.ProductID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(MessageResponse{
				Message: "Failed to delete cart item",
			})
		}
		return c.JSON(MessageResponse{Message: "Cart item deleted"})
	}

	if err := h.Repo.ModifyItem(body.Username, body.ProductID, body.Quantity, body.UnitPrice); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(MessageResponse{
			Message: "Failed to update cart",
		})
	}

	return c.JSON(MessageResponse{Message: "Cart updated"})
}

// GetCartItems returns the user's cart items
// @Summary      Get cart items
// @Description  Fetches all cart items associated with a username
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        username query string true "Username"
// @Success      200 {array} models.CartItem
// @Failure      400 {object} handlers.MessageResponse
// @Failure      500 {object} handlers.MessageResponse
// @Router       /cart/items [get]
func (h *CartHandler) GetCartItems(c *fiber.Ctx) error {
	username := c.Query("username")

	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(MessageResponse{
			Message: "Username is required",
		})
	}

	items, err := h.Repo.GetItemsByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(MessageResponse{
			Message: "Could not retrieve cart items",
		})
	}

	return c.JSON(items)
}
