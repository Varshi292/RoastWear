package bootstrap

import (
	"github.com/Varshi292/RoastWear/internal/database"
	"github.com/Varshi292/RoastWear/internal/handlers"
	"github.com/Varshi292/RoastWear/internal/repositories"
	"github.com/Varshi292/RoastWear/internal/services"
	"github.com/Varshi292/RoastWear/internal/sessions"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
)

func InitializeApp() (*fiber.App, string) {
	// Load .env
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
	app := initializeFiber(appCfg)

	// Static files
	app.Static("/", "./frontend/build")
	app.Static("/uploads", "./uploads")

	// Database
	userSqliteDB := database.NewSqliteUserDatabase(appCfg.UserDBPath)
	userDB, err := userSqliteDB.Connect()
	if err != nil {
		log.Fatalf("User database connection error: %s", err)
	}
	log.Println("✅ Connected to user database successfully")
	if err := userSqliteDB.Migrate(); err != nil {
		log.Fatalf("User database migration error: %s", err)
	}
	log.Println("✅ Migrated user database successfully")

	sessionSqliteDB := database.NewSqliteSessionDatabase(appCfg.SessionDBPath)
	sessionDB, err := sessionSqliteDB.Connect()
	if err != nil {
		log.Fatalf("Session database connection error: %s", err)
	}
	log.Println("✅ Connected to session database successfully")
	if err := sessionSqliteDB.Migrate(); err != nil {
		log.Fatalf("Session database migration error: %s", err)
	}
	log.Println("✅ Migrated session database successfully")

	uploadSqliteDB := database.NewSqliteUploadDatabase(appCfg.UploadDBPath)
	uploadDB, err := uploadSqliteDB.Connect()
	if err != nil {
		log.Fatalf("Upload database connection error: %s", err)
	}
	log.Println("✅ Connected to upload database successfully")
	if err := uploadSqliteDB.Migrate(); err != nil {
		log.Fatalf("Upload database migration error: %s", err)
	}
	log.Println("✅ Migrated upload database successfully")

	// Dependencies
	userRepo := &repositories.UserRepository{Db: userDB}
	userService := services.NewUserService(userRepo)
	sessionService := repositories.NewSessionRepository(sessionDB)
	authService := services.NewAuthService(userRepo, sessionService)

	// Handlers
	registerHandler := handlers.NewRegisterHandler(userService)
	loginHandler := handlers.NewLoginHandler(authService, sessionService)
	sessionHandler := handlers.NewSessionHandler(sessionService)
	checkoutHandler := handlers.NewCheckoutHandler(sessionService, sessionDB)

	// Routes
	app.Post("/register", registerHandler.UserRegister)
	app.Post("/login", loginHandler.UserLogin)

	app.Post("/session/create", sessionHandler.CreateSession)
	app.Post("/session/verify", sessionHandler.VerifySession)
	app.Delete("/session/delete", sessionHandler.DeleteSession)

	app.Post("/checkout", checkoutHandler.CheckoutCart)
	app.Post("/post_user_image", handlers.UploadImageHandler(uploadDB))
	app.Get("/get_user_images", handlers.GetImagesHandler(uploadDB))

	app.Get("/docs/*", swagger.HandlerDefault)

	return app, ":" + appCfg.BackendPort
}
