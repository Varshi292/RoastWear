package main

import (
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/repository"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/session"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Load environmental variables
	cfg := config.AppConfig{}
	if err := config.LoadConfig(&cfg); err != nil {
		log.Fatalf("Failed to load app configuration: %s", err)
	}
	if err := config.LoadConfig(config.SessionConfig{}); err != nil {
		log.Fatalf("Failed to load session configuration: %s", err)
	}
	sessCfg := config.SessionConfig{}
	session.InitializeSessionStore(sessCfg)

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
	registerHandler := handlers.NewRegisterHandler(userService)
	app.Post("/register", registerHandler.UserRegister)

	authService := services.NewAuthService(userRepo)
	loginHandler := handlers.NewLoginHandler(authService)
	app.Post("/login", loginHandler.UserLogin)

	log.Printf("Server running on port %s\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
