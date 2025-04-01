package main

import (
	"log"

	_ "github.com/Varshi292/RoastWear/docs"
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/repository"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

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
		AllowOrigins:     "http://localhost:7777, http://127.0.0.1:7777",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization",
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

	// Repositories
	userRepo := &repository.UserRepository{Db: db}

	// Services
	userService := services.NewUserService(userRepo)
	sessionService := services.NewSessionService(db)
	authService := services.NewAuthService(userRepo, sessionService)

	// Handlers
	registerHandler := handlers.NewRegisterHandler(userService)
	loginHandler := handlers.NewLoginHandler(authService, sessionService)
	sessionHandler := handlers.NewSessionHandler(sessionService)
	checkoutHandler := handlers.NewCheckoutHandler(sessionService, db) // ✅ NEW

	// Routes
	app.Post("/register", registerHandler.UserRegister)
	app.Post("/login", loginHandler.UserLogin)

	app.Post("/session/create", sessionHandler.CreateSession)
	app.Post("/session/verify", sessionHandler.VerifySession)
	app.Delete("/session/delete", sessionHandler.DeleteSession)

	app.Post("/checkout", checkoutHandler.CheckoutCart) // ✅ NEW: handles cart submission

	app.Post("/post_user_image", handlers.UploadImageHandler(db))
	app.Get("/get_user_images", handlers.GetImagesHandler(db))

	// Docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Start the server at configured port
	log.Printf("Server running on port %s\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
