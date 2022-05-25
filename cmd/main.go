package main

import (
	"github.com/ValikoDorodnov/passport/internal/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	for _, route := range router.DefineRoutes() {
		http.HandleFunc(route.Path, route.Handler)
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
