// internal/handlers/purchase_handler.go
package handlers

import (
	"fmt"
	"net/http"

	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type PurchaseHandler struct {
	CartRepo     *repositories.CartRepository
	PurchaseRepo *repositories.PurchaseRepository
}

func NewPurchaseHandler(cartRepo *repositories.CartRepository, purchaseRepo *repositories.PurchaseRepository) *PurchaseHandler {
	return &PurchaseHandler{
		CartRepo:     cartRepo,
		PurchaseRepo: purchaseRepo,
	}
}

func (h *PurchaseHandler) Checkout(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
	}

	if err := c.BodyParser(&body); err != nil || body.Username == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or missing username"})
	}

	// Step 1: Fetch user's cart items
	items, err := h.CartRepo.GetItemsByUsername(body.Username)
	if err != nil {
		fmt.Println("❌ Failed to retrieve cart:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cart"})
	}
	if len(items) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cart is empty"})
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
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save purchase"})
	}

	// Step 4: Clear the cart
	if err := h.CartRepo.ClearCartForUser(body.Username); err != nil {
		fmt.Println("⚠️ Cart not cleared:", err)
		// Still return 200, but notify of partial success
		return c.Status(http.StatusOK).JSON(fiber.Map{"success": true, "message": "Purchase saved, but cart not cleared"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Purchase completed and cart cleared"})
}
