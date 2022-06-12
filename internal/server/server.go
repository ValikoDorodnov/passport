package server

import "database/sql"

type Env struct {
	Conn *sql.DB
}

func NewEnv(db *sql.DB) *Env {
	return &Env{
		Conn: db,
	}
}
