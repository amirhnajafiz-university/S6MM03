package cmd

import (
	"log"

	"github.com/amirhnajafiz/hollyworld/back-end/internal/handler"
	"github.com/gofiber/fiber/v2"
)

// Execute our main application
func Execute() {
	// creating a fiber app
	app := fiber.New()

	// creating a handler
	h := handler.Handler{}

	// defining our endpoints
	app.Get("/api/movies", h.GetTopMovies)
	app.Get("/api/movies/:id", h.GetSingleMovie)
	app.Get("/dash/:id", h.GetMovieFile)

	// listening on port 8080
	log.Fatal(app.Listen(":8080"))
	// TODO: migrate file
}
