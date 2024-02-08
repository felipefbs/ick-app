package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/websocket"
)

func Init(db *sql.DB) *http.Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(hotReload)
	// router.Use(VerifyCookieSession)

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

func hotReload(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != "/reload" {
			next.ServeHTTP(w, r)
			return
		}

		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			err := websocket.Message.Send(ws, "reload")
			if err != nil {
				fmt.Println(err)
			}

			err = websocket.Message.Receive(ws, nil)
			if err != nil {
				fmt.Println(err)
			}
		}).ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
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
