package models

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	ID        int64  `json:"id" db:"id" example:"1"`
	AuthorID  uuid.UUID  `json:"author_id" db:"author_id" example:"1"`
	Ttl       int    `json:"ttl" db:"ttl" example:"300"`
	Price     int    `json:"price" db:"price"  example:"300"`
	SetAt  time.Time `json:"set_at" db:"set_at" example:"0000-01-01T00:00:00.000000Z"`
}