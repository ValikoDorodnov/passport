package router

import (
	"fmt"
	"net/http"
)

func DefineRoutes() []RouteStruct {
	var routes []RouteStruct

	routes = append(routes, RouteStruct{"/", handleMain})

	return routes
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello main")
}
