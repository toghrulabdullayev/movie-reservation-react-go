package showtimes

import (
	"time"

	"github.com/google/uuid"
)

// showtimes
type ShowtimeStatus string

const (
	Scheduled ShowtimeStatus = "scheduled"
	Ongoing   ShowtimeStatus = "ongoing"
	Completed ShowtimeStatus = "completed"
	Canceled  ShowtimeStatus = "cancelled"
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

// showtime seats
type ShowtimeSeatStatus string

const (
	SeatAvailable ShowtimeSeatStatus = "available"
	SeatHeld      ShowtimeSeatStatus = "held"
	SeatBooked    ShowtimeSeatStatus = "booked"
	SeatBlocked   ShowtimeSeatStatus = "blocked"
)

type ShowtimeSeat struct {
	ID         uuid.UUID          `json:"id"`
	ShowtimeID uuid.UUID          `json:"showtime_id"`
	SeatID     uuid.UUID          `json:"seat_id"`
	Status     ShowtimeSeatStatus `json:"status"`
	Price      float64            `json:"price"` // base_price * multiplier
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}
