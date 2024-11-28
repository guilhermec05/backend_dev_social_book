package router

import (
	"database/sql"
	router "dev_book_backend/internal/http/router/Routers"

	"github.com/gorilla/mux"
)

func ConfigRouter(db *sql.DB) *mux.Router {

	r := mux.NewRouter()

	return router.Config(r, db)
}
