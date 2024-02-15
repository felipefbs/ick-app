package handlers

import (
	"log/slog"
	"net/http"

	"github.com/felipefbs/ick-app/internal/helpers"
	"github.com/felipefbs/ick-app/internal/repositories"
	"github.com/felipefbs/ick-app/templates"
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

func (handler *IckHandler) IckListPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, _ := helpers.IsLogged(ctx)

	ickList, err := handler.repo.Get(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find general ick list", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	userIckList, err := handler.repo.FindUserIcks(ctx, userID)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find user ick list", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = templates.IckListPage(true, ickList, userID.String(), userIckList).Render(ctx, w)
	if err != nil {
		slog.ErrorContext(ctx, "failed to render ick list page", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (handler *IckHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	_, isLogged := helpers.IsLogged(r.Context())

	err := templates.MainPage(isLogged).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *IckHandler) DefinitionPage(w http.ResponseWriter, r *http.Request) {
	// err := templates.Main(templates.Definition()).Render(r.Context(), w)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
}

func (handler *IckHandler) RegisterIckPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, isLogged := helpers.IsLogged(ctx)

	ickList, err := handler.repo.Get(ctx, userID)

	err = templates.RegisterIck(isLogged, ickList).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *IckHandler) RegisterIck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, _ := helpers.IsLogged(ctx)
	ickValue := r.FormValue("ick")

	err := handler.repo.Save(ctx, ickValue, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	ickList, err := handler.repo.Get(ctx, userID)

	err = templates.IckList(ickList).Render(ctx, w)
	if err != nil {
		slog.Error("failed to render template", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *IckHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	// coo, err := r.Cookie("session-cookie")
	// if err != nil {
	// 	slog.Error("failed to get cookie", "error", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// userID, err := uuid.Parse(coo.Value)
	// if err != nil {
	// 	return
	// }

	// ickID, err := uuid.Parse(chi.URLParam(r, "ick-id"))
	// if err != nil {
	// 	slog.Error("invalid ick id", "error", err, "id", chi.URLParam(r, "ick-id"))
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// err = handler.repo.Upvote(r.Context(), userID, ickID)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
}
