package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DbFile = "./sqlite.db"
)

// NewConnection opens a connection to our sqlite file
func NewConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", DbFile)
}

// Migrate our database file
func Migrate(conn *sql.DB) error {
	_, err := conn.Exec(CreateTableQuery)

	return err
}

// Seed data into our database
func Seed(conn *sql.DB) error {
	type Movie struct {
		Title       string
		Director    string
		Score       int
		Description string
		Poster      string
		Link        string
	}

	s, err := conn.Prepare(InsertMovieQuery)
	if err != nil {
		return err
	}

	movies := []Movie{
		{
			Title:       "test",
			Director:    "test",
			Score:       2,
			Description: "test",
			Poster:      "test",
			Link:        "test",
		},
	}

	for movie := range movies {
		_, err = s.Exec(movie)
		if err != nil {
			return err
		}
	}

	return nil
}
