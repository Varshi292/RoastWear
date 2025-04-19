package main

import (
	"github.com/Varshi292/RoastWear/internal/bootstrap"
	"github.com/gofiber/fiber/v2/log"
)

// @title RoastWear API
// @version 1.0
// @description This is the backend API for the RoastWear application.
// @host localhost:7777
// @BasePath /

// main Entry point for the RoastWear application
func main() {
	app, port := bootstrap.InitializeApp()
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
