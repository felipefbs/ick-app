package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(db *pgxpool.Pool) *http.Server {
	router := chi.NewRouter()

	registerIckRoutes(router, db)
	registerUserRoutes(router, db)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return server
}
