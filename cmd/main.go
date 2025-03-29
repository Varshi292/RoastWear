package main

import (
	"github.com/Varshi292/RoastWear/internal/configurations"
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/models"
	"github.com/Varshi292/RoastWear/internal/repository"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Load environmental variables
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	// Initialize Fiber
	app := fiber.New()
	app.Static("/", cfg.StaticFilesPath)

	// Initialize database
	db, err := database.Open(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	userRepo := &repository.UserRepository{Db: db}
	userService := services.NewUserService(userRepo)
	if err := userService.RegisterUser(&models.UserCreateRequest{
		Username: "hi",
		Email:    "hi",
		Password: "hi",
	}); err != nil {
		log.Fatalf("Failed to register user: %s", err)
	}

	log.Printf("Server running on port %s\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
