package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init(db *sql.DB) *http.Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	registerIckRoutes(router, db)
	registerUserRoutes(router, db)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return server
}
