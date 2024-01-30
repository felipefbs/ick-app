package icks

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/felipefbs/ick-app/templates"
	"github.com/google/uuid"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (handler *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Main().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *Handler) IckPage(w http.ResponseWriter, r *http.Request) {
	err := templates.RegisterIck().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *Handler) DefinitionPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Definition().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *Handler) RegisterIck(w http.ResponseWriter, r *http.Request) {

	ick := r.FormValue("ick")
	coo, err := r.Cookie("session-cookie")
	if err != nil {
		log.Println(err)
	}

	userID := uuid.UUID{}
	if coo.Valid() == nil {
		err = handler.db.QueryRow("select id from users where username = $1", coo.Value).Scan(&userID)
		if err != nil {
			log.Println(err)
		}
	}

	err = templates.RegisterIck().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = handler.db.Exec("INSERT INTO icks (id, ick, registered_by) values ($1, $2, $3)", uuid.New(), ick, userID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
