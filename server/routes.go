package server

import (
	"database/sql"

	"github.com/felipefbs/ick-app/ick"
	"github.com/felipefbs/ick-app/user"
	"github.com/go-chi/chi/v5"
)

func registerIckRoutes(router chi.Router, db *sql.DB) {
	repo := ick.NewRepository(db)
	userRepo := user.NewRepository(db)
	handler := ick.NewHandler(repo, userRepo)

	router.Get("/", handler.MainPage)
	router.Get("/ick-list", handler.ListPage)
	router.Get("/definition", handler.DefinitionPage)
	router.Get("/register", handler.IckPage)

	router.Put("/upvote/{ick-id}", handler.Upvote)

	router.Post("/register", handler.RegisterIck)
}

func registerUserRoutes(router chi.Router, db *sql.DB) {
	repo := user.NewRepository(db)
	handler := user.NewHandler(repo)

	router.Get("/register-user", handler.RegisterPage)
	router.Post("/login", handler.Login)
	router.Post("/logout", handler.Logout)
	router.Post("/register-user", handler.RegisterUser)
}
