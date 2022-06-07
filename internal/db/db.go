package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

var Connection *pgx.Conn

func Init() {
	// следует закрывать соединение в месте вызова
	// defer conn.Close(context.Background())

	var err error
	Connection, err = pgx.Connect(context.Background(), os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

	}
}
