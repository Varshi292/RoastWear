package bootstrap

import (
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/swagger"
)

func InitializeApp() (*fiber.App, string) {
	// Load .env
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load app config: %s", err)
	}

	// Load app config
	appCfg, err := loadAppConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to load app config: %s", err)
	}

	// Load session config
	sessCfg, err := loadSessionConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to load session config: %s", err)
	}

	session.InitializeSessionStore(sessCfg)

	// Fiber setup
	app := initializeFiber(appCfg)

	// Static files
	app.Static("/", "./frontend/build")
	app.Static("/uploads", "./uploads")

	initializeAdmin(app)

	// Database
	s := database.NewSqliteDatabase(appCfg.DBPath)
	db, err := s.Connect()
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
	if err := s.Migrate(); err != nil {
		log.Fatalf("DB migration error: %s", err)
	}

	// Dependencies
	userRepo := &repositories.UserRepository{Db: db}
	userService := services.NewUserService(userRepo)
	sessionService := services.NewSessionService(db)
	authService := services.NewAuthService(userRepo, sessionService)

	// Handlers
	registerHandler := handlers.NewRegisterHandler(userService)
	loginHandler := handlers.NewLoginHandler(authService, sessionService)
	sessionHandler := handlers.NewSessionHandler(sessionService)
	checkoutHandler := handlers.NewCheckoutHandler(sessionService, db)

	// Routes
	app.Post("/register", registerHandler.UserRegister)
	app.Post("/login", loginHandler.UserLogin)

	app.Post("/session/create", sessionHandler.CreateSession)
	app.Post("/session/verify", sessionHandler.VerifySession)
	app.Delete("/session/delete", sessionHandler.DeleteSession)

	app.Post("/checkout", checkoutHandler.CheckoutCart)
	app.Post("/post_user_image", handlers.UploadImageHandler(db))
	app.Get("/get_user_images", handlers.GetImagesHandler(db))

	app.Get("/docs/*", swagger.HandlerDefault)

	return app, ":" + appCfg.BackendPort
}
