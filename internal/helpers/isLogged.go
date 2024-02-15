package helpers

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func IsLogged(ctx context.Context) (uuid.UUID, bool) {
	isLogged, ok := ctx.Value("isLogged").(bool)
	if !ok {
		slog.ErrorContext(ctx, "failed to verify user session")

		return uuid.Nil, false
	}

	if !isLogged {
		slog.InfoContext(ctx, "user is not logged")

		return uuid.Nil, false
	}

	user, ok := ctx.Value("user").(string)
	if !ok {
		slog.ErrorContext(ctx, "failed to cast user id")

		return uuid.Nil, false
	}

	userID, err := uuid.Parse(user)
	if err != nil {
		slog.ErrorContext(ctx, "invalid user id", "user", user, "error", err)
	}

	return userID, isLogged
}
