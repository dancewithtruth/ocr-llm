package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id        uuid.UUID
	IPAddress string
	CreatedAt time.Time
	UpdatedAt time.Time
}
