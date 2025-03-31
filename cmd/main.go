package main

import (
	_ "github.com/Varshi292/RoastWear/docs"
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/repository"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
)

import "github.com/gofiber/fiber/v2/middleware/cors"

// @title RoastWear API
// @version 1.0
// @description This is the backend API for the RoastWear application.
// @host localhost:7777
// @BasePath /

// main Entry point for the RoastWear application. Manages configurations,
// services, repositories, routes, and handlers.
func main() {
	// Load environmental variables
	cfg := config.AppConfig{}
	if err := config.LoadConfig(&cfg); err != nil {
		log.Fatalf("Failed to load app configuration: %s", err)
	}

	// Initialize session store
	sessCfg := config.SessionConfig{}
	if err := config.LoadConfig(sessCfg); err != nil {
		log.Fatalf("Failed to load session configuration: %s", err)
	}
	session.InitializeSessionStore(sessCfg)

	// Initialize Fiber
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:7777, http://localhost:7777",
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))
	

	// Retrieve frontend build for deployment
	app.Static("/", cfg.StaticFilesPath)

	// Initialize and migrate the database
	db, err := database.Open(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	// Set up user repository
	userRepo := &repository.UserRepository{Db: db}
	// Set up user service
	userService := services.NewUserService(userRepo)
	// Set up authentication service
	authService := services.NewAuthService(userRepo)

	// Set up register and login handlers
	registerHandler := handlers.NewRegisterHandler(userService)
	app.Post("/register", registerHandler.UserRegister)
	loginHandler := handlers.NewLoginHandler(authService)
	app.Post("/login", loginHandler.UserLogin)

	app.Post("/post_user_image", handlers.UploadImageHandler(db))
	app.Get("/get_user_images", handlers.GetImagesHandler(db))


	// Serve docs at /docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Start the server at configured port
	log.Printf("Server running on port %s\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
