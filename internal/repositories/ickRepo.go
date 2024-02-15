package repositories

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/felipefbs/ick-app/pkg/ick"
	"github.com/felipefbs/ick-app/pkg/user"
	"github.com/google/uuid"
)

type IckRepository struct {
	db *sql.DB
}

func NewIckRepository(db *sql.DB) *IckRepository {
	return &IckRepository{db: db}
}

func (repo *IckRepository) Save(ctx context.Context, ick string, userID uuid.UUID) error {
	if ick == "" {
		slog.Warn("cant save empty ick")

		return nil
	}

	tx, err := repo.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		slog.Error("failed to start transaction", "error", err)

		return err
	}

	defer tx.Rollback()

	ickID := uuid.New()
	_, err = tx.ExecContext(ctx,
		"INSERT INTO icks (id, ick, registered_by) values (?, ?, ?)",
		ickID, ick, userID)
	if err != nil {
		slog.Error("failed to save ick into database", "error", err, "table", "icks")
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO user_icks (user_id, icks_id) values (?, ?)", userID, ickID)
	if err != nil {
		slog.Error("failed to save ick into database", "error", err, "table", "user_icks")
		return err
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("failed to commit transaction", "error", err)

		return err
	}

	return nil
}

func (repo *IckRepository) Upvote(ctx context.Context, userID, ickID uuid.UUID) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO user_icks (user_id, icks_id) values (?, ?)", userID, ickID)
	if err != nil {
		slog.Error("failed to save ick into database", "error", err, "table", "user_icks")
		return err
	}

	return nil
}

func (repo *IckRepository) Get(ctx context.Context, userID ...uuid.UUID) ([]ick.Ick, error) {
	tx, err := repo.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})
	if err != nil {
		slog.Error("failed to start transaction", "error", err)

		return nil, err
	}

	defer tx.Rollback()

	var rows *sql.Rows

	if len(userID) == 0 {
		rows, err = tx.QueryContext(ctx, "SELECT id, ick, registered_by FROM icks order by created_at desc")
		if err != nil {
			slog.Error("failed to list all icks", "error", err, "table", "icks")

			return nil, err
		}
	} else {
		rows, err = tx.QueryContext(ctx, "SELECT id, ick, registered_by FROM icks where registered_by = ? order by created_at desc", userID[0])
		if err != nil {
			slog.Error("failed to list all icks", "error", err, "table", "icks")

			return nil, err
		}
	}

	defer rows.Close()

	ickList := make([]ick.Ick, 0)
	for rows.Next() {
		var id, registeredBy uuid.UUID
		var ickName string

		err := rows.Scan(&id, &ickName, &registeredBy)
		if err != nil {
			slog.Error("failed to scan an ick", "error", err)
		}

		foundIck := ick.Ick{ID: id, Ick: ickName, RegisteredBy: registeredBy}

		var username string
		err = tx.QueryRowContext(ctx, "SELECT username FROM users where id = ?", registeredBy).Scan(&username)
		if err != nil {
			slog.Error("failed to scan an ick", "error", err, "table", "users")
		}

		foundIck.User = user.User{
			ID:       registeredBy,
			Username: username,
		}

		ickList = append(ickList, foundIck)
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("failed to commit transaction", "error", err)

		return nil, err
	}

	return ickList, nil
}

func (repo *IckRepository) FindUserIcks(ctx context.Context, userID uuid.UUID) (map[uuid.UUID]bool, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT icks_id from user_icks where user_id = ?", userID)
	if err != nil {
		slog.Error("failed to get users icks", "error", err, "table", "user_icks")

		return nil, err
	}

	defer rows.Close()

	userIckList := make(map[uuid.UUID]bool)

	for rows.Next() {
		var ickID uuid.UUID
		err := rows.Scan(&ickID)
		if err != nil {
			slog.Error("failed to scan ick id", "error", err, "user id", userID)
		}

		userIckList[ickID] = true
	}

	return userIckList, nil
}
