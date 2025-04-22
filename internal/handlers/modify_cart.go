package handlers

import (
	"fmt"

	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	Repo *repositories.CartRepository
}

func NewCartHandler(repo *repositories.CartRepository) *CartHandler {
	return &CartHandler{Repo: repo}
}

func (h *CartHandler) ModifyCart(c *fiber.Ctx) error {
	fmt.Println("📥 Received request to /cart/modify")

	var body struct {
		Username  string  `json:"username"`
		ProductID int     `json:"productid"`
		Quantity  int     `json:"quantity"`
		UnitPrice float64 `json:"unitPrice"`
	}

	bodyRaw := c.Body()
	fmt.Println("🧾 Raw body:", string(bodyRaw))

	if err := c.BodyParser(&body); err != nil {
		fmt.Println("❌ Failed to parse body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if body.Username == "" {
		fmt.Println("❌ Missing username in request")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You must be logged in to store cart information."})
	}

	fmt.Printf("🧑 Username: %s\n", body.Username)
	fmt.Printf("🛒 ProductID: %d\n", body.ProductID)
	fmt.Printf("🔢 Quantity: %d\n", body.Quantity)
	fmt.Printf("💲 TotalPrice: %.2f\n", body.UnitPrice)

	// If quantity is zero, delete the row
	if body.Quantity == 0 {
		err := h.Repo.DeleteItem(body.Username, body.ProductID)
		if err != nil {
			fmt.Println("❌ Error deleting cart item in DB:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cart item"})
		}
		fmt.Println("✅ Cart item deleted successfully")
		return c.JSON(fiber.Map{"message": "Cart item deleted"})
	}

	if err := h.Repo.ModifyItem(body.Username, body.ProductID, body.Quantity, body.UnitPrice); err != nil {
		fmt.Println("❌ Error updating cart in DB:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update cart"})
	}

	fmt.Println("✅ Cart modification successful")
	return c.JSON(fiber.Map{"message": "Cart updated"})
}

func (h *CartHandler) GetCartItems(c *fiber.Ctx) error {
	username := c.Query("username") // grab from query param

	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username is required"})
	}

	items, err := h.Repo.GetItemsByUsername(username)
	if err != nil {
		fmt.Println("❌ Failed to get cart items:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve cart items"})
	}

	return c.JSON(items)
}
