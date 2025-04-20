package bootstrap

import (
	"github.com/Varshi292/RoastWear/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func initializeFiber(config *config.AppConfig) *fiber.App {
	app := fiber.New()
	x := cors.Config{
		AllowOrigins:     config.Domain + ":" + config.FrontendPort,
		AllowCredentials: true,
		AllowMethods:     "GET,POST,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization",
	}
	app.Use(cors.New(x))
	return app
}
