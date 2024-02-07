package handlers

import (
	"log/slog"
	"net/http"

	"github.com/felipefbs/ick-app/internal/repositories"
	"github.com/felipefbs/ick-app/templates"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type IckHandler struct {
	repo     *repositories.IckRepository
	userRepo *repositories.UserRepository
}

func NewIckHandler(repo *repositories.IckRepository, userRepo *repositories.UserRepository) *IckHandler {
	return &IckHandler{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (handler *IckHandler) ListPage(w http.ResponseWriter, r *http.Request) {
	ickList, err := handler.repo.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	coo, err := r.Cookie("session-cookie")
	if err != nil {
		slog.Error("failed to get cookie", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(coo.Value)
	if err != nil {
		slog.Error("failed to parse user id", "error", err)
	}

	userIckList, err := handler.repo.FindUserIcks(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = templates.Main(templates.IckList(ickList, coo.Value, userIckList)).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *IckHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Main(templates.Definition()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *IckHandler) IckPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Main(templates.RegisterIck()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *IckHandler) DefinitionPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Main(templates.Definition()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *IckHandler) RegisterIck(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie("session-cookie")
	if err != nil {
		slog.Error("failed to get session cookie", "error", err)
	}

	userID := uuid.UUID{}
	if err := coo.Valid(); err == nil {
		userID, _ = uuid.Parse(coo.Value)
	}

	ick := r.FormValue("ick")
	slog.Info(ick)

	err = handler.repo.Save(r.Context(), ick, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = templates.Main(templates.RegisterIck()).Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render template", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *IckHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie("session-cookie")
	if err != nil {
		slog.Error("failed to get cookie", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(coo.Value)
	if err != nil {
		return
	}

	ickID, err := uuid.Parse(chi.URLParam(r, "ick-id"))
	if err != nil {
		slog.Error("invalid ick id", "error", err, "id", chi.URLParam(r, "ick-id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.repo.Upvote(r.Context(), userID, ickID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}