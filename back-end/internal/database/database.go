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
