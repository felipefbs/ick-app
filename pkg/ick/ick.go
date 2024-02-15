package ick

import (
	"github.com/felipefbs/ick-app/pkg/user"
	"github.com/google/uuid"
)

type Ick struct {
	ID           uuid.UUID
	Ick          string
	Upvotes      int
	Downvotes    int
	RegisteredBy uuid.UUID
	User         user.User
}
