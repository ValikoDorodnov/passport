package controller

import (
	"fmt"
	"github.com/ValikoDorodnov/passport/internal/app/article"
	"github.com/ValikoDorodnov/passport/internal/server"
	"net/http"
)

func GetArticle(env *server.Env, w http.ResponseWriter, r *http.Request) error {
	defer env.Conn.Close()

	var a article.Article
	sql := "SELECT id, name, text FROM article WHERE id = 1"

	err := env.Conn.QueryRow(sql).Scan(&a.Id, &a.Name, &a.Text)

	fmt.Fprintf(w, "Article: %+v", a)
	return err
}
