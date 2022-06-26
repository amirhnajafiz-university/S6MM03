package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// Handler is our endpoint handling struct
type Handler struct {
	Db *sql.Conn
}

// GetTopMovies returns 6 top movies
func (h *Handler) GetTopMovies(c *fiber.Ctx) error {
	return nil
}

// GetSingleMovie returns a single movie information by id
func (h *Handler) GetSingleMovie(c *fiber.Ctx) error {
	return nil
}

// GetMovieFile returns a movie thriller file
func (h *Handler) GetMovieFile(c *fiber.Ctx) error {
	return nil
}
