package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/felipefbs/ick-app/entities"
	"github.com/felipefbs/ick-app/templates"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (handler *Handler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	err := templates.RegisterUser().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (handler *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	coo := &http.Cookie{
		Name:     "session-cookie",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, coo)
	err := templates.RegisterUser().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	loginInfo := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.repo.GetByUsername(r.Context(), loginInfo.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	coo := &http.Cookie{
		Name:     "session-cookie",
		Value:    user.ID.String(),
		MaxAge:   3600 * 24,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, coo)
	err = templates.RegisterIck().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user, err := entities.NewUserFromRequestBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = handler.repo.Save(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	coo := &http.Cookie{
		Name:     "session-cookie",
		Value:    user.ID.String(),
		MaxAge:   3600 * 24,
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(w, coo)

	err = templates.RegisterIck().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
