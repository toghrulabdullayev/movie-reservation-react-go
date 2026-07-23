package rooms

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Capacity  int       `json:"capacity"`
	Rows      int       `json:"rows"`
	Columns   int       `json:"columns"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
