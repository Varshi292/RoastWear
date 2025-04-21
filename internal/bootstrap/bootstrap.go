package bootstrap

import (
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
)

func InitializeApp() (*fiber.App, string) {
	// Initialize config validator
	if err := utils.InitializeValidator(); err != nil {
		log.Fatalf("Failed to initialize validator: %s", err)
	}

	// Load .env config
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load app config: %s", err)
	}
	log.Printf("✅ .env configuration loaded successfully.")

	// Load app config
	appCfg, err := loadAppConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to load app config: %s", err)
	}
	log.Printf("✅ App configuration loaded successfully.")

	// Load session config
	sessCfg, err := loadSessionConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to load session config: %s", err)
	}
	log.Printf("✅ Session configuration loaded successfully.")

	sessions.InitializeSessionStore(sessCfg)

	// Fiber setup
	app := InitializeFiber(appCfg)

	// Static files
	app.Static("/", "./frontend/build")
	app.Static("/uploads", "./uploads")

	// Databases
	userDB := InitializeDatabase(database.NewSqliteUserDatabase(appCfg.UserDBPath))
	sessionDB := InitializeDatabase(database.NewSqliteSessionDatabase(appCfg.SessionDBPath))
	uploadDB := InitializeDatabase(database.NewSqliteUploadDatabase(appCfg.UploadDBPath))
	log.Println("✅ Databases initialized successfully")

	// Dependencies
	userRepo := &repositories.UserRepository{Db: userDB}
	userService := services.NewUserService(userRepo)
	sessionRepo := repositories.NewSessionRepository(sessionDB)
	authService := services.NewAuthService(userRepo, sessionRepo)
	log.Println("✅ Dependencies initialized successfully")

	// Handlers
	registerHandler := handlers.NewRegisterHandler(userService, sessionRepo)
	loginHandler := handlers.NewLoginHandler(authService, sessionRepo)
	sessionHandler := handlers.NewSessionHandler(sessionRepo)
	logoutHandler := handlers.NewLogoutHandler(sessionRepo)
	checkoutHandler := handlers.NewCheckoutHandler(sessionRepo, sessionDB)

	// Routes
	app.Post("/register", registerHandler.UserRegister)
	app.Post("/login", loginHandler.UserLogin)
	app.Delete("/logout", logoutHandler.UserLogout)

	app.Post("/session/create", sessionHandler.CreateSession)
	app.Get("/session/verify", sessionHandler.VerifySession)
	app.Delete("/session/delete", sessionHandler.DeleteSession)

	app.Post("/checkout", checkoutHandler.CheckoutCart)
	app.Post("/post_user_image", handlers.UploadImageHandler(uploadDB))
	app.Get("/get_user_images", handlers.GetImagesHandler(uploadDB))

	app.Get("/docs/*", swagger.HandlerDefault)

	return app, ":" + appCfg.BackendPort
}
