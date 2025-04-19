package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func initializeFiber() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:7777, http://127.0.0.1:7777",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization",
	}))
	return app
}
