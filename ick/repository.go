package ick

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Save(ctx context.Context, ick string, userID uuid.UUID) error {
	_, err := repo.db.ExecContext(ctx,
		"INSERT INTO icks (id, ick, registered_by) values (?, ?, ?)",
		uuid.New(), ick, userID)
	if err != nil {
		slog.Error("failed to save ick into database", "error", err)
		return err
	}

	return nil
}
