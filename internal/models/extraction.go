package models

import (
	"time"

	"github.com/google/uuid"
)

type Extraction struct {
	Id        uuid.UUID
	SessionId uuid.UUID
	Result    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
