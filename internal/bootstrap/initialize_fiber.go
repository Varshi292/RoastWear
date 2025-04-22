package bootstrap

import (
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func initializeFiber(config *config.AppConfig) *fiber.App {
	app := fiber.New()
	x := cors.Config{
		AllowOrigins:     "http://127.0.0.1:7777", // must match browser origin exactly
		AllowCredentials: true,
		AllowMethods:     "GET,POST,OPTIONS,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
	}
	app.Use(cors.New(x))
	return app
}
