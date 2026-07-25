package reservations

import (
	"time"

	"github.com/google/uuid"
)

type ReservationResponse struct {
	ID         uuid.UUID         `json:"id"`
	UserID     uuid.UUID         `json:"user_id"`
	ShowtimeID uuid.UUID         `json:"showtime_id"`
	Status     ReservationStatus `json:"status"`
	TotalPrice float64           `json:"total_price"`
	ReservedAt time.Time         `json:"reserved_at"`
}

type CreateReservationRequest struct {
	UserID     uuid.UUID `json:"user_id"`
	ShowtimeID uuid.UUID `json:"showtime_id"`
	TotalPrice float64   `json:"total_price"`
}

type UpdateReservationRequest struct {
	UserID     uuid.UUID         `json:"user_id"`
	ShowtimeID uuid.UUID         `json:"showtime_id"`
	Status     ReservationStatus `json:"status"`
	TotalPrice float64           `json:"total_price"`
}

func toReservationResponse(r *Reservation) ReservationResponse {
	if r == nil {
		return ReservationResponse{}
	}

	return ReservationResponse{
		ID:         r.ID,
		UserID:     r.UserID,
		ShowtimeID: r.ShowtimeID,
		Status:     r.Status,
		TotalPrice: r.TotalPrice,
		ReservedAt: r.ReservedAt,
	}
}
