package entities

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Birthdate string    `json:"birthdate"`
	Password  string    `json:"password"`
}

func NewUser(username, name, gender, birthdate, password string) (*User, error) {
	byteArray, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to generate hashed password", "error", err)
		return nil, err
	}

	return &User{
		Username:  username,
		Name:      name,
		Gender:    gender,
		Birthdate: birthdate,
		Password:  string(byteArray),
	}, nil
}

func NewUserFromRequestBody(body io.ReadCloser) (*User, error) {
	user := User{}
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		slog.Error("failed to decode request body", "error", err)
		return nil, err
	}

	body.Close()

	return NewUser(user.Username, user.Name, user.Gender, user.Birthdate, user.Password)
}
