package handler

import (
	"database/sql"
	"log"
	"path"

	"github.com/gofiber/fiber/v2"
)

// BASE folder of Dash files
const (
	BASE   = "/tmp/dash"
	POSTER = "./storage/files"
)

// Handler is our endpoint handling struct
type Handler struct {
	Db *sql.DB
}

// GetTopMovies returns top 6 movies
func (h *Handler) GetTopMovies(c *fiber.Ctx) error {
	// creating a movie type
	type Movie struct {
		Uid   int    `json:"id"`
		Title string `json:"title"`
	}

	// creating our variables
	var (
		movie  Movie
		movies []Movie

		query = "SELECT uid, title FROM movies ORDER BY score LIMIT 6"
	)

	// set cors policy
	c.Set("Access-Control-Allow-Origin", "localhost")
	c.Set("Content-Type", "application/json")

	// executing database query
	rows, err := h.Db.Query(query)
	if err != nil {
		return err
	}

	// scan output
	for rows.Next() {
		er := rows.Scan(
			&movie.Uid,
			&movie.Title,
		)
		if er != nil {
			return er
		}

		movies = append(movies, movie)
	}

	log.Println("GET TOP")

	// close rows
	_ = rows.Close()

	return c.JSON(movies)
}

// GetSingleMovie returns a single movie information by id
func (h *Handler) GetSingleMovie(c *fiber.Ctx) error {
	// creating a movie type
	type Movie struct {
		Uid         int    `json:"id"`
		Title       string `json:"title"`
		Director    string `json:"director"`
		Score       int    `json:"score"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Link        string `json:"link"`
	}

	// creating our variables
	var (
		movie Movie

		query = "SELECT * FROM movies WHERE uid=?"
	)

	// set cors policy
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Content-Type", "application/json")

	// creating a prepare
	s, _ := h.Db.Prepare(query)

	// executing query
	rows, err := s.Query(c.Params("id"))
	if err != nil {
		return err
	}

	// scan into movie
	for rows.Next() {
		err = rows.Scan(
			&movie.Uid,
			&movie.Title,
			&movie.Director,
			&movie.Score,
			&movie.Description,
			&movie.Type,
			&movie.Link,
		)
		if err != nil {
			return err
		}
	}

	log.Printf("GET %s\n", c.Params("id"))

	return c.JSON(movie)
}

// GetMoviePoster returns the movie poster
func (h *Handler) GetMoviePoster(c *fiber.Ctx) error {
	// set cors policy
	c.Set("Access-Control-Allow-Origin", "*")

	log.Printf("GET poster by ID %s\n", c.Params("id"))

	return c.SendFile(path.Join(POSTER, c.Params("id")+".png"))
}

// GetMovieFile returns a movie thriller file
func (h *Handler) GetMovieFile(c *fiber.Ctx) error {
	// set cors policy
	c.Set("Access-Control-Allow-Origin", "*")

	log.Printf("GET file by ID %s\n", c.Params("id"))

	return c.SendFile(path.Join(BASE, c.Params("id")+".mpd"))
}
