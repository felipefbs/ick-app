package server

import (
	"database/sql"

	"github.com/felipefbs/ick-app/internal/handlers"
	"github.com/felipefbs/ick-app/internal/repositories"
	"github.com/go-chi/chi/v5"
)

func registerIckRoutes(router chi.Router, db *sql.DB) {
	repo := repositories.NewIckRepository(db)
	userRepo := repositories.NewUserRepository(db)
	handler := handlers.NewIckHandler(repo, userRepo)

	router.Get("/", handler.MainPage)

	router.Get("/ick-list", handler.ListPage)
	router.Get("/definition", handler.DefinitionPage)
	router.Get("/register", handler.IckPage)

	router.Put("/upvote/{ick-id}", handler.Upvote)

	router.Post("/register", handler.RegisterIck)
}

func registerUserRoutes(router chi.Router, db *sql.DB) {
	repo := repositories.NewUserRepository(db)
	handler := handlers.NewUserHandler(repo)

	router.Get("/register-user", handler.RegisterPage)
	router.Post("/login", handler.Login)
	router.Post("/logout", handler.Logout)
	router.Post("/register-user", handler.RegisterUser)
}
