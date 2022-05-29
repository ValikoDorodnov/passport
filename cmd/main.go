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

	handler := router.New()

	if err := http.ListenAndServe(":8099", handler); err != nil {
		log.Fatal(err)
	}
}
