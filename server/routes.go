package server

import (
	"database/sql"

	"github.com/felipefbs/ick-app/icks"
	"github.com/felipefbs/ick-app/user"
	"github.com/go-chi/chi/v5"
)

func registerIckRoutes(router chi.Router, db *sql.DB) {
	handler := icks.NewHandler(db)

	router.Get("/", handler.MainPage)
	router.Get("/definition", handler.DefinitionPage)
	router.Get("/register", handler.IckPage)

	router.Post("/register", handler.RegisterIck)
}

func registerUserRoutes(router chi.Router, db *sql.DB) {
	handler := user.NewHandler(db)

	router.Get("/register-user", handler.RegisterPage)
	router.Post("/login", handler.Login)
	router.Post("/logout", handler.Logout)
	router.Post("/register-user", handler.RegisterUser)
}
