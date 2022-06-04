package router

import (
	"encoding/json"
	"fmt"
	"github.com/ValikoDorodnov/passport/internal/app/article"
	"net/http"
	"regexp"
	"strings"
)

type User struct {
	Test string
}

type router struct {
}

type route struct {
	Method  string
	regex   *regexp.Regexp
	Handler func(w http.ResponseWriter, r *http.Request)
}

var routes = []route{
	newRoute("GET", "/", handleMain),
	newRoute("POST", "/auth", handleAuth),
	newRoute("GET", "/article", handleArticle),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello main")
}

func handleAuth(w http.ResponseWriter, r *http.Request) {

	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User: %+v", u)
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	a := article.GetArticle()

	fmt.Fprintf(w, "Article: %+v", a)
}

func (router router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var nowAllow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.Method {
				nowAllow = append(nowAllow, route.Method)
				continue
			}
			route.Handler(w, r)
			return
		}
	}
	if len(nowAllow) > 0 {
		w.Header().Set("Allow", strings.Join(nowAllow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

func New() *router {
	return &router{}
}
