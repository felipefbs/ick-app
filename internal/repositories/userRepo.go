package repositories

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/felipefbs/ick-app/pkg/user"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Save(ctx context.Context, user *user.User) error {
	_, err := repo.db.ExecContext(ctx,
		"INSERT INTO users (id, username, name, birthdate, gender, password) values ($1, $2, $3, $4, $5, $6)",
		user.ID, user.Username, user.Name, user.Birthdate, user.Gender, user.Password)
	if err != nil {
		slog.Error("failed to save user into database", "error", err)

		return err
	}

	return nil
}

func (repo *UserRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	user := user.User{}
	err := repo.db.QueryRow("SELECT id, username, password from users where username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		slog.Error("failed to find user", "error", err, "username", username)

		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	user := user.User{}

	err := repo.db.QueryRow("SELECT id, username, name, gender, birthdate from users where id = ?", userID).Scan(&user.ID, &user.Username, &user.Name, &user.Gender, &user.Birthdate)
	if err != nil {
		slog.Error("failed to find user", "error", err, "user id", userID)

		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetUserIDByUsername(ctx context.Context, username string) (uuid.UUID, error) {
	id := uuid.UUID{}
	err := repo.db.QueryRowContext(ctx, "select id from users where username = ?", username).Scan(&id)
	if err != nil {
		slog.Error("failed to find user", "error", err, "username", username)

		return uuid.Nil, err
	}

	return id, nil
}
