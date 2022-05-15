package main

import (
	"github.com/ValikoDorodnov/passport/internal/router"
	"log"
	"net/http"
)

func main() {
	for _, route := range router.DefineRoutes() {
		http.HandleFunc(route.Path, route.Handler)
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
