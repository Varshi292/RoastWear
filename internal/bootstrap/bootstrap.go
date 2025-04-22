package bootstrap

import (
	"log"

	_ "github.com/Varshi292/RoastWear/docs"
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/Varshi292/RoastWear/internal/utils"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

// InitializeApp initializes and sets up the entire application, including configuration, services, routes, and databases.
// @Title RoastWear API
// @Description This is the API for RoastWear, a system for user management, session handling, and image uploads.
// @Version 1.0
// @BasePath /
// @Schemes http
// @Consumes application/json
// @Produces application/json
// @Tags Auth, Cart, Session, Images
func InitializeApp() (*fiber.App, string) {
	// Initialize config validator
	if err := utils.InitializeValidator(); err != nil {
		log.Fatalf("Failed to initialize validator: %s", err)
	}

	// Load .env config
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load .env config: %s", err)
	}
	log.Println("✅ .env configuration loaded successfully.")

	appCfg, err := loadAppConfig(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to load app config: %s", err)
	}
	log.Println("✅ App configuration loaded successfully.")

	sessCfg, err := loadSessionConfig(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to load session config: %s", err)
	}
	log.Println("✅ Session configuration loaded successfully.")
	sessions.InitializeSessionStore(sessCfg)

	// Initialize Fiber
	app := InitializeFiber(appCfg)

	// Serve static files
	app.Static("/", "./frontend/build")
	app.Static("/uploads", "./uploads")
	app.Static("/static", "./frontend/build/static") // Serve /static for JS/CSS assets

	// Initialize databases
	userDB := InitializeDatabase(database.NewSqliteUserDatabase(appCfg.UserDBPath))
	sessionDB := InitializeDatabase(database.NewSqliteSessionDatabase(appCfg.SessionDBPath))
	uploadDB := InitializeDatabase(database.NewSqliteUploadDatabase(appCfg.UploadDBPath))
	cartDB := database.NewSqliteCartDatabase(appCfg.CartDBPath)
	purchaseDB := database.NewSqlitePurchaseDatabase(appCfg.CartDBPath) // temporarily reusing CartDBPath

	cartConn, err := cartDB.Connect()
	if err != nil {
		log.Fatalf("❌ Failed to connect to cart DB: %v", err)
	}
	if err := cartDB.Migrate(); err != nil {
		log.Fatalf("❌ Failed to migrate cart DB: %v", err)
	}

	purchaseConn, err := purchaseDB.Connect()
	if err != nil {
		log.Fatalf("❌ Failed to connect to purchase DB: %v", err)
	}
	if err := purchaseDB.Migrate(); err != nil {
		log.Fatalf("❌ Failed to migrate purchase DB: %v", err)
	}

	log.Println("✅ All databases initialized successfully.")

	// Initialize repositories and services
	userRepo := &repositories.UserRepository{Db: userDB}
	sessionRepo := repositories.NewSessionRepository(sessionDB)
	authService := services.NewAuthService(userRepo, sessionRepo)
	userService := services.NewUserService(userRepo)
	cartRepo := repositories.NewCartRepository(cartConn)
	purchaseRepo := repositories.NewPurchaseRepository(purchaseConn)

	// Initialize handlers
	registerHandler := handlers.NewRegisterHandler(userService, sessionRepo)
	loginHandler := handlers.NewLoginHandler(authService, sessionRepo)
	sessionHandler := handlers.NewSessionHandler(sessionRepo)
	logoutHandler := handlers.NewLogoutHandler(sessionRepo)
	cartHandler := handlers.NewCartHandler(cartRepo)
	purchaseHandler := handlers.NewPurchaseHandler(cartRepo, purchaseRepo)

	// Register routes
	app.Post("/register", registerHandler.UserRegister)
	app.Post("/login", loginHandler.UserLogin)
	app.Delete("/logout", logoutHandler.UserLogout)

	app.Post("/session/create", sessionHandler.CreateSession)
	app.Get("/session/verify", sessionHandler.VerifySession)
	app.Delete("/session/delete", sessionHandler.DeleteSession)

	app.Post("/post_user_image", handlers.UploadImageHandler(uploadDB))
	app.Get("/get_user_images", handlers.GetImagesHandler(uploadDB))

	app.Post("/cart/modify", cartHandler.ModifyCart)
	app.Post("/checkout", purchaseHandler.Checkout)

	app.Get("/docs/*", swagger.WrapHandler)

	app.Get("/cart/items", cartHandler.GetCartItems)

	return app, ":" + appCfg.BackendPort
}
