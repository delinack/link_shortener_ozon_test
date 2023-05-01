package model

import (
	"github.com/google/uuid"
	"time"
)

type Link struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Token     string    `json:"token" db:"token"`
	URL       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
