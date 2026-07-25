package seats

import (
	"time"

	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID `json:"id"`
	RoomID     uuid.UUID `json:"room_id"`
	SeatTypeID uuid.UUID `json:"seat_type_id"`
	Row        int       `json:"row"`
	Number     int       `json:"number"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SeatType struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"` // "STANDARD", "VIP", etc.
	Multiplier float64   `json:"multiplier"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
