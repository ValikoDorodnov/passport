package router

import "net/http"

type RouteStruct struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}
