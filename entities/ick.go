package entities

import (
	"github.com/google/uuid"
)

type Ick struct {
	ID           uuid.UUID
	Ick          string
	RegisteredBy uuid.UUID
	User         User
}
