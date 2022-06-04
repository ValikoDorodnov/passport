package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func Connection() *pgx.Conn {
	// следует закрывать соединение в месте вызова
	// defer conn.Close(context.Background())

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

	}

	return conn
}
