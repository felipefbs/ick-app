package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Init(db *sql.DB) *http.Server {
	router := chi.NewRouter()

	registerIckRoutes(router, db)
	registerUserRoutes(router, db)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return server
}
