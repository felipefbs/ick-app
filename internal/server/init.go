package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *http.Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(VerifyCookieSession)

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	registerUserRoutes(router)
	registerIckRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	slog.Info("initialing http server", "addr", server.Addr)

	return server
}

func VerifyCookieSession(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		coo, err := r.Cookie("session-cookie")
		if err != nil || coo.Valid() != nil {
			ctx := context.WithValue(r.Context(), "isLogged", false)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

			return
		}

		ctx := context.WithValue(r.Context(), "user", coo.Value)
		ctx = context.WithValue(ctx, "isLogged", true)
		slog.InfoContext(ctx, "user info", "userID", coo.Value)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
