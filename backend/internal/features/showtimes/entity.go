package showtimes

import (
	"time"

	"github.com/google/uuid"
)

type ShowtimeStatus string

const (
	Scheduled ShowtimeStatus = "scheduled"
	Ongoing   ShowtimeStatus = "ongoing"
	Completed ShowtimeStatus = "completed"
	Canceled  ShowtimeStatus = "canceled"
)

type Showtime struct {
	ID        uuid.UUID      `json:"id"`
	MovieID   uuid.UUID      `json:"movie_id"`
	RoomID    uuid.UUID      `json:"room_id"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	BasePrice float64        `json:"base_price"`
	Status    ShowtimeStatus `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
