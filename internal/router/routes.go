package router

import (
	"fmt"
	"github.com/ValikoDorodnov/passport/internal/app/jwt"
	"net/http"
)

func DefineRoutes() []RouteStruct {
	var routes []RouteStruct

	routes = append(routes, RouteStruct{"/", handleMain})
	routes = append(routes, RouteStruct{"/build-token", handleBuildToken})
	routes = append(routes, RouteStruct{"/parse-token", handleParseToken})

	return routes
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello main")
}

func handleBuildToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.BuildToken()

	fmt.Fprint(w, string(token))
}

func handleParseToken(w http.ResponseWriter, r *http.Request) {
	parsed := jwt.ParseToken([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NTM0OTg3MjMsImlzcyI6InNvbWUtaXNzdWVyIn0.kyYzusECEy_Q2qxKLKBdxPfKwd0cw35QQGHnhvyS_YQ"))

	fmt.Fprint(w, parsed.Issuer())
}
