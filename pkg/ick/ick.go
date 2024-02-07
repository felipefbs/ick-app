package ick

import (
	"github.com/felipefbs/ick-app/pkg/user"
	"github.com/google/uuid"
)

type Ick struct {
	ID           uuid.UUID
	Ick          string
	RegisteredBy uuid.UUID
	User         user.User
}
