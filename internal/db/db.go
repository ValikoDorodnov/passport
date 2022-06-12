package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func Init() (*sql.DB, error) {

	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))

	if err != nil {
		return nil, err
	}

	return db, err
}
