package models

import "time"

type User struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
