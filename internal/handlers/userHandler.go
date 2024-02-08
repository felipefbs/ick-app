package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/felipefbs/ick-app/internal/helpers"
	"github.com/felipefbs/ick-app/internal/repositories"
	"github.com/felipefbs/ick-app/pkg/user"
	"github.com/felipefbs/ick-app/templates"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (handler *UserHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	err := templates.LoginPage().Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render template", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginInfo := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		slog.Error("failed to decode login info", "error", err)

		w.Header().Set("HX-Retarget", "#login-form")
		w.WriteHeader(http.StatusOK) //ToDO: Make it right -> this is so wrong

		templates.LoginComponent(false).Render(r.Context(), w)

		return
	}

	if loginInfo.Username == "" || loginInfo.Password == "" {
		slog.Error("empty username or password", "username", loginInfo.Username)

		w.Header().Set("HX-Retarget", "#login-form")
		w.WriteHeader(http.StatusOK) //ToDO: Make it right -> this is so wrong

		templates.LoginComponent(true).Render(r.Context(), w)
		return
	}

	user, err := handler.repo.GetByUsername(r.Context(), loginInfo.Username)
	if err != nil {
		slog.Error("failed to find the user in database", "error", err, "username", loginInfo.Username)

		w.Header().Set("HX-Retarget", "#login-form")
		w.WriteHeader(http.StatusOK) //ToDO: Make it right -> this is so wrong

		templates.LoginComponent(true).Render(r.Context(), w)

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		slog.Error("invalid username or password", "error", err, "user", user.ID)

		w.WriteHeader(http.StatusOK) //ToDO: Make it right -> this is so wrong
		templates.LoginComponent(true).Render(r.Context(), w)

		return
	}

	helpers.SetSessionCookie(w, user.ID.String())
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func (handler *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	coo := &http.Cookie{
		Name:     "session-cookie",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, coo)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func (handler *UserHandler) RegisterUserPage(w http.ResponseWriter, r *http.Request) {
	err := templates.RegisterPage().Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render template", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user, err := user.NewUserFromRequestBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = handler.repo.Save(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	helpers.SetSessionCookie(w, user.ID.String())
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func (handler *UserHandler) ProfilePage(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie("session-cookie")
	if err != nil {
		slog.Error("failed to get session cookie", "error", err)
		return
	}

	userID, err := uuid.Parse(coo.Value)
	if err != nil {
		slog.Error("failed to parse user id", "error", err, "id", coo.Value)
	}

	user, err := handler.repo.GetByUserID(r.Context(), userID)
	if err != nil {
		slog.Error("failed to find user", "error", err, "id", userID)
	}

	err = templates.ProfilePage(user).Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render template", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
