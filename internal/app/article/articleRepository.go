package article

import (
	"context"
	"fmt"
	"github.com/ValikoDorodnov/passport/internal/db"
	"os"
)

func GetArticle() Article {
	conn := db.Connection()
	defer conn.Close(context.Background())

	var article Article

	sql := "SELECT id, name, text FROM article WHERE id = 1"
	err := conn.QueryRow(context.Background(), sql).Scan(&article.id, &article.name, &article.text)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return article
}
