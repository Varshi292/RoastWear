package main

import (
	"github.com/Varshi292/RoastWear/internal/bootstrap"
	"log"
)

// main Entry point for the RoastWear application
// @title RoastWear API
// @version 1.0
// @description This is the backend API for the RoastWear application.
// @host localhost:7777
// @BasePath /
func main() {
	app, port := bootstrap.InitializeApp()
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
