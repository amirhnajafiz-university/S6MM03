package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Execute our main application
func Execute() {
	// creating a fiber app
	app := fiber.New()

	// listening on port 8080
	log.Fatal(app.Listen(":8080"))
	// TODO: fiber handler
	// TODO: migrate file
	// TODO: top 6 endpoint
	// TODO: a single endpoint
	// TODO: serve file
}
