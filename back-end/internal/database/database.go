package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DbFile = "./sqlite.db"
)

// NewConnection opens a connection to our sqlite file
func NewConnection() (*sql.Conn, error) {
	db, err := sql.Open("sqlite3", DbFile)
	if err != nil {
		return nil, err
	}

	return db.Conn(context.Background())
}
