package user

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/felipefbs/ick-app/entities"
	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) Save(ctx context.Context, user *entities.User) error {
	_, err := repo.db.ExecContext(ctx,
		"INSERT INTO users (id, username, name, birthdate, gender, password) values ($1, $2, $3, $4, $5, $6)",
		uuid.New(), user.Username, user.Name, user.Birthdate, user.Gender, user.Password)
	if err != nil {
		slog.Error("failed to save user into database", "error", err)

		return err
	}

	return nil
}

func (repo *Repository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.QueryRow("SELECT id, username, password from users where username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		slog.Error("failed to find user", "error", err, "username", username)

		return nil, err
	}

	return &user, nil
}

func (repo *Repository) GetUserIDByUsername(ctx context.Context, username string) (uuid.UUID, error) {
	id := uuid.UUID{}
	err := repo.db.QueryRowContext(ctx, "select id from users where username = ?", username).Scan(&id)
	if err != nil {
		slog.Error("failed to find user", "error", err, "username", username)

		return uuid.Nil, err
	}

	return id, nil
}
