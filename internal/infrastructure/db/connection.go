package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_CONNECTION_STRING")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
