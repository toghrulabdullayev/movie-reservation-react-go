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

type ShowtimeSeatStatus string

const (
	SeatAvailable ShowtimeSeatStatus = "AVAILABLE"
	SeatHeld      ShowtimeSeatStatus = "HELD"
	SeatBooked    ShowtimeSeatStatus = "BOOKED"
	SeatBlocked   ShowtimeSeatStatus = "BLOCKED"
)

type ShowtimeSeat struct {
	ID         uuid.UUID          `json:"id"`
	ShowtimeID uuid.UUID          `json:"showtime_id"`
	SeatID     uuid.UUID          `json:"seat_id"`
	Status     ShowtimeSeatStatus `json:"status"`
	Price      float64            `json:"price"` // calculated base_price * multiplier
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}
