package main

import (
	"RoastWear/internal/admin/models"
	"RoastWear/internal/admin/pages"
	"RoastWear/internal/config"
	"RoastWear/internal/database"
	"RoastWear/internal/handlers"
	"RoastWear/internal/repositories"
	"RoastWear/internal/services"
	"RoastWear/internal/session"
	_ "github.com/GoAdminGroup/go-admin/adapter/gofiber"
	"github.com/GoAdminGroup/go-admin/engine"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/tests/tables"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
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

	template.AddComp(chartjs.NewChart())
	eng := engine.Default()
	if err := eng.AddConfigFromYAML("./internal/admin/config.yml").
		AddGenerators(tables.Generators).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})
	models.Init(eng.SqliteConnection())
	app.Static("/uploads", "./uploads")

	// Retrieve frontend build for deployment
	app.Static("/", cfg.StaticFilesPath)

	// Initialize and migrate the database
	s := database.NewSqliteDatabase(cfg.DBPath)
	db, err := s.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	if err := s.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	// Set up user repository
	userRepo := &repositories.UserRepository{Db: db}
	// Set up user service
	userService := services.NewUserService(userRepo)
	// Set up authentication service
	authService := services.NewAuthService(userRepo)

	// Set up register and login handlers
	registerHandler := handlers.NewRegisterHandler(userService)
	app.Post("/register", registerHandler.UserRegister)
	loginHandler := handlers.NewLoginHandler(authService)
	app.Post("/login", loginHandler.UserLogin)

	// Serve docs at /docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Start the server at configured port
	log.Printf("Server running on port %s\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
	eng.SqliteConnection().Close()
}
