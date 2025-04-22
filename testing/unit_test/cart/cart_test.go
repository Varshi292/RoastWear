package cart_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupTestApp(t *testing.T) (*fiber.App, *repositories.CartRepository) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}
	db.AutoMigrate(&models.CartItem{})
	repo := repositories.NewCartRepository(db)
	handler := handlers.NewCartHandler(repo)

	app := fiber.New()
	app.Post("/cart/modify", handler.ModifyCart)
	app.Get("/cart/items", func(c *fiber.Ctx) error {
		username := c.Query("username")
		items, err := repo.GetItemsByUsername(username)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch items"})
		}
		return c.JSON(items)
	})
	return app, repo
}

func TestCart_AddItems_NotLoggedIn(t *testing.T) {
	app, _ := setupTestApp(t)

	req := httptest.NewRequest(http.MethodPost, "/cart/modify", bytes.NewBuffer([]byte(`{
		"username": "",
		"productid": 1,
		"quantity": 2,
		"unitPrice": 19.99
	}`)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil || resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized, got %d", resp.StatusCode)
	}
}

func TestCart_LoginAndSyncCart(t *testing.T) {
	_, repo := setupTestApp(t)
	repo.ModifyItem("guest", 1, 2, 19.99)
	repo.ModifyItem("guest", 2, 1, 24.99)

	guestItems, _ := repo.GetItemsByUsername("guest")
	for _, item := range guestItems {
		price, err := strconv.ParseFloat(item.TotalPrice, 64)
		if err != nil {
			t.Errorf("Failed to parse total price: %v", err)
			continue
		}
		repo.ModifyItem("grillio", item.ProductID, item.Quantity, price)
	}
	repo.ClearCartForUser("guest")

	updated, _ := repo.GetItemsByUsername("grillio")
	if len(updated) != 2 {
		t.Errorf("Expected 2 items transferred, got %d", len(updated))
	}
}

func TestCart_PurchaseAfterAdding(t *testing.T) {
	_, repo := setupTestApp(t)
	repo.ModifyItem("grillio", 1, 2, 19.99)
	repo.ModifyItem("grillio", 2, 1, 24.99)
	repo.ClearCartForUser("grillio")
	items, _ := repo.GetItemsByUsername("grillio")
	if len(items) != 0 {
		t.Errorf("Expected cart to be empty after purchase, found %d items", len(items))
	}
}

func TestCart_PurchaseWithoutLogin(t *testing.T) {
	_, repo := setupTestApp(t)
	err := repo.ClearCartForUser("")
	if err == nil {
		t.Error("Expected error or no-op when trying to purchase without login")
	}
}

func TestCart_LoggedInAddThenPurchase(t *testing.T) {
	_, repo := setupTestApp(t)
	repo.ModifyItem("grillio", 1, 3, 19.99)
	items, _ := repo.GetItemsByUsername("grillio")
	if len(items) != 1 || items[0].Quantity != 3 {
		t.Errorf("Expected 1 item with quantity 3, got %v", items)
	}
	repo.ClearCartForUser("grillio")
	empty, _ := repo.GetItemsByUsername("grillio")
	if len(empty) != 0 {
		t.Errorf("Expected cart to be cleared, found %d items", len(empty))
	}
}

func TestCart_AddThenDeleteAllThenAttemptPurchase(t *testing.T) {
	_, repo := setupTestApp(t)
	repo.ModifyItem("grillio", 1, 1, 19.99)
	repo.ModifyItem("grillio", 2, 2, 24.99)
	repo.ModifyItem("grillio", 1, 0, 0)
	repo.ModifyItem("grillio", 2, 0, 0)
	items, _ := repo.GetItemsByUsername("grillio")
	if len(items) != 0 {
		t.Errorf("Expected cart to be empty after deletions, found %d items", len(items))
	}
	err := repo.ClearCartForUser("grillio")
	if err != nil {
		t.Errorf("Unexpected error when clearing empty cart: %v", err)
	}
}
