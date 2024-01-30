package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/felipefbs/ick-app/templates"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
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

	user := User{}
	err = handler.db.QueryRow("SELECT id, username, password from users where username=$1", loginInfo.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Println(err)
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
		Value:    user.Username,
		MaxAge:   3600,
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
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = NewUser(user.Username, user.Name, user.Gender, user.Birthdate, user.Password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = handler.db.Exec(
		"INSERT INTO users (id, username, name, birthdate, gender, password) values ($1, $2, $3, $4, $5, $6)",
		uuid.New(), user.Username, user.Name, user.Birthdate, user.Gender, user.Password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	coo := &http.Cookie{
		Name:     "session-cookie",
		Value:    user.Username,
		MaxAge:   3600,
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
