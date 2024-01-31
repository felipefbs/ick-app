package ick

import (
	"log/slog"
	"net/http"

	"github.com/felipefbs/ick-app/templates"
	"github.com/felipefbs/ick-app/user"
	"github.com/google/uuid"
)

type Handler struct {
	repo     *Repository
	userRepo *user.Repository
}

func NewHandler(repo *Repository, userRepo *user.Repository) *Handler {
	return &Handler{
		repo:     repo,
		userRepo: userRepo,
	}
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
	coo, err := r.Cookie("session-cookie")
	if err != nil {
		slog.Error("failed to get session cookie", "error", err)
	}

	userID := uuid.UUID{}
	if err := coo.Valid(); err == nil {
		userID, _ = handler.userRepo.GetUserIDByUsername(r.Context(), coo.Value)
	}

	ick := r.FormValue("ick")

	err = handler.repo.Save(r.Context(), ick, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = templates.RegisterIck().Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render template", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
