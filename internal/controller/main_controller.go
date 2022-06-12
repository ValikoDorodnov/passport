package controller

import (
	"fmt"
	"github.com/ValikoDorodnov/passport/internal/server"
	"net/http"
)

func GetIndex(env *server.Env, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "Hello main")
	return nil
}
