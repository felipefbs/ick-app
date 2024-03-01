package server

import (
	"log/slog"
	"net/http"

	"github.com/felipefbs/ick-app/internal/database"
	"github.com/felipefbs/ick-app/internal/handlers"
	"github.com/felipefbs/ick-app/internal/repositories"
	"github.com/go-chi/chi/v5"
)

func registerIckRoutes(router chi.Router) {
	db, err := database.Get()
	if err != nil {
		slog.Error("failed to get database", "error", err)

		return
	}

	repo := repositories.NewIckRepository(db)
	userRepo := repositories.NewUserRepository(db)
	handler := handlers.NewIckHandler(repo, userRepo)

	router.Handle("/", http.RedirectHandler("/home", http.StatusPermanentRedirect))
	router.Get("/home", handler.MainPage)

	router.Get("/ick-list", handler.IckListPage)

	router.Get("/ick", handler.RegisterIckPage)
	router.Post("/ick", handler.RegisterIck)
}

func registerUserRoutes(router chi.Router) {
	db, err := database.Get()
	if err != nil {
		slog.Error("failed to get database", "error", err)

		return
	}

	repo := repositories.NewUserRepository(db)
	handler := handlers.NewUserHandler(repo)

	router.Get("/profile", handler.ProfilePage)
	router.Post("/logout", handler.Logout)

	router.Get("/login", handler.LoginPage)
	router.Post("/login", handler.Login)

	router.Get("/register-user", handler.RegisterUserPage)
	router.Post("/register-user", handler.RegisterUser)

}
