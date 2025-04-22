package handlers

import (
	"fmt"
	"net/http"

	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

// CheckoutRequest represents the expected body for a purchase checkout
type CheckoutRequest struct {
	Username string `json:"username"`
}

// PurchaseResponse represents a generic success response
type PurchaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

// PurchaseHandler handles purchases and checkout
type PurchaseHandler struct {
	CartRepo     *repositories.CartRepository
	PurchaseRepo *repositories.PurchaseRepository
}

// NewPurchaseHandler creates a new handler for purchases
func NewPurchaseHandler(cartRepo *repositories.CartRepository, purchaseRepo *repositories.PurchaseRepository) *PurchaseHandler {
	return &PurchaseHandler{
		CartRepo:     cartRepo,
		PurchaseRepo: purchaseRepo,
	}
}

// Checkout handles the checkout process for a user.
// @Summary      Checkout and complete purchase
// @Description  Processes the items in the user's cart, creates a purchase record, and clears the cart.
// @Tags         purchase
// @Accept       json
// @Produce      json
// @Param        checkoutRequest body handlers.CheckoutRequest true "Checkout request with username"
// @Success      200 {object} handlers.PurchaseResponse "Purchase completed and cart cleared"
// @Failure      400 {object} handlers.ErrorResponse "Invalid input or empty cart"
// @Failure      500 {object} handlers.ErrorResponse "Internal server error during purchase processing"
// @Router       /checkout [post]
func (h *PurchaseHandler) Checkout(c *fiber.Ctx) error {
	var body CheckoutRequest

	if err := c.BodyParser(&body); err != nil || body.Username == "" {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid or missing username",
		})
	}

	// Step 1: Fetch user's cart items
	items, err := h.CartRepo.GetItemsByUsername(body.Username)
	if err != nil {
		fmt.Println("❌ Failed to retrieve cart:", err)
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
			Error:   "Failed to retrieve cart",
			Details: err.Error(),
		})
	}
	if len(items) == 0 {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
			Error: "Cart is empty",
		})
	}

	// Step 2: Format product string and calculate total
	var productsStr string
	var total float64
	for i, item := range items {
		if i > 0 {
			productsStr += "#"
		}
		productsStr += fmt.Sprintf("%d:%d", item.ProductID, item.Quantity)
		var price float64
		fmt.Sscanf(item.TotalPrice, "%f", &price)
		total += price
	}

	// Step 3: Save the purchase record
	err = h.PurchaseRepo.CreatePurchase(body.Username, productsStr, fmt.Sprintf("%.2f", total))
	if err != nil {
		fmt.Println("❌ Failed to record purchase:", err)
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
			Error:   "Failed to save purchase",
			Details: err.Error(),
		})
	}

	// Step 4: Clear the cart
	if err := h.CartRepo.ClearCartForUser(body.Username); err != nil {
		fmt.Println("⚠️ Cart not cleared:", err)
		return c.Status(http.StatusOK).JSON(PurchaseResponse{
			Success: true,
			Message: "Purchase saved, but cart not cleared",
		})
	}

	return c.JSON(PurchaseResponse{
		Success: true,
		Message: "Purchase completed and cart cleared",
	})
}
