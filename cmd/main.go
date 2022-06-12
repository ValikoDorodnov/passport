package main

import (
	"github.com/ValikoDorodnov/passport/internal/controller"
	"github.com/ValikoDorodnov/passport/internal/db"
	"github.com/ValikoDorodnov/passport/internal/handler"
	"github.com/ValikoDorodnov/passport/internal/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	env := server.NewEnv(conn)
	http.Handle("/", handler.Handler{env, controller.GetIndex})
	http.Handle("/article", handler.Handler{env, controller.GetArticle})

	log.Fatal(http.ListenAndServe(":8099", nil))
}
