package reservations

import (
	"time"

	"github.com/google/uuid"
)

type ReservationStatus string

const (
	ReservationPending   ReservationStatus = "pending"
	ReservationConfirmed ReservationStatus = "confirmed"
	ReservationCancelled ReservationStatus = "cancelled"
)

type Reservation struct {
	ID         uuid.UUID         `json:"id"`
	UserID     uuid.UUID         `json:"user_id"`
	ShowtimeID uuid.UUID         `json:"showtime_id"`
	Status     ReservationStatus `json:"status"`
	TotalPrice float64           `json:"total_price"`
	ReservedAt time.Time         `json:"reserved_at"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// ReservationSeat acts as the bridge model matching reservation_seats
type ReservationSeat struct {
	ReservationID  uuid.UUID `json:"reservation_id"`
	ShowtimeSeatID uuid.UUID `json:"showtime_seat_id"`
}
