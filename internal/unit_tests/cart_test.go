package repositories_test

import (
	"testing"

	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}
	db.AutoMigrate(&models.CartItem{})
	return db
}

func TestCart_AddItems_NotLoggedIn(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewCartRepository(db)

	err := repo.ModifyItem("", 1, 2, 19.99)
	if err == nil {
		t.Error("Expected error when modifying cart with empty username")
	}
}

func TestCart_LoginAndSyncCart(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewCartRepository(db)

	repo.ModifyItem("guest", 1, 2, 19.99)
	repo.ModifyItem("guest", 2, 1, 24.99)

	// simulate login and transfer to user
	guestItems, _ := repo.GetItemsByUsername("guest")
	for _, item := range guestItems {
		repo.ModifyItem("grillio", item.ProductID, item.Quantity, 19.99)
	}
	repo.ClearCartForUser("guest")

	updated, _ := repo.GetItemsByUsername("grillio")
	if len(updated) != 2 {
		t.Errorf("Expected 2 items transferred, got %d", len(updated))
	}
}

func TestCart_PurchaseAfterAdding(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewCartRepository(db)

	repo.ModifyItem("grillio", 1, 2, 19.99)
	repo.ModifyItem("grillio", 2, 1, 24.99)

	repo.ClearCartForUser("grillio")
	items, _ := repo.GetItemsByUsername("grillio")
	if len(items) != 0 {
		t.Errorf("Expected cart to be empty after purchase, found %d items", len(items))
	}
}

func TestCart_PurchaseWithoutLogin(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewCartRepository(db)

	err := repo.ClearCartForUser("")
	if err == nil {
		t.Error("Expected error or no-op when trying to purchase without login")
	}
}

func TestCart_LoggedInAddThenPurchase(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewCartRepository(db)

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
	db := setupTestDB(t)
	repo := repositories.NewCartRepository(db)

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
