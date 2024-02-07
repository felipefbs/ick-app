package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init(db *sql.DB) *http.Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(VerifyCookieSession)

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	registerUserRoutes(router, db)
	registerIckRoutes(router, db)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return server
}

func VerifyCookieSession(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/register-user" || r.URL.Path == "/login" {
			fmt.Println(r.URL.Path)
			next.ServeHTTP(w, r)

			return
		}

		coo, err := r.Cookie("session-cookie")
		if err != nil || coo.Valid() != nil {
			http.Redirect(w, r, "/register-user", http.StatusPermanentRedirect)
			next.ServeHTTP(w, r)

			return
		}

		ctx := context.WithValue(r.Context(), "user", coo.Value)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
