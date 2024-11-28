package router

import (
	"database/sql"
	"dev_book_backend/internal/http/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Config(r *mux.Router, db *sql.DB) *mux.Router {
	var routers []Router
	routers = append(routers, NewControllerUser(db)...)
	routers = append(routers, NewFollowRotes(db)...)
	routers = append(routers, NewPublishRotes(db)...)

	for _, router := range routers {

		if router.RequireAuth {
			r.HandleFunc(router.URI, middleware.Log(middleware.Autentication(router.Function))).Methods(router.Method)
		} else {

			r.HandleFunc(router.URI, middleware.Log(router.Function)).Methods(router.Method)
		}

	}

	return r
}
