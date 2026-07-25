package showtimes

import (
	"time"

	"github.com/google/uuid"
)

type ShowtimeResponse struct {
	ID        uuid.UUID      `json:"id"`
	MovieID   uuid.UUID      `json:"movie_id"`
	RoomID    uuid.UUID      `json:"room_id"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	BasePrice float64        `json:"base_price"`
	Status    ShowtimeStatus `json:"status"`
}

type CreateShowtimeRequest struct {
	MovieID   uuid.UUID      `json:"movie_id"`
	RoomID    uuid.UUID      `json:"room_id"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	BasePrice float64        `json:"base_price"`
	Status    ShowtimeStatus `json:"status"`
}

type UpdateShowtimeRequest struct {
	MovieID   uuid.UUID      `json:"movie_id"`
	RoomID    uuid.UUID      `json:"room_id"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	BasePrice float64        `json:"base_price"`
	Status    ShowtimeStatus `json:"status"`
}

type ShowtimeSeatResponse struct {
	ID         uuid.UUID          `json:"id"`
	ShowtimeID uuid.UUID          `json:"showtime_id"`
	SeatID     uuid.UUID          `json:"seat_id"`
	Status     ShowtimeSeatStatus `json:"status"`
	Price      float64            `json:"price"` // calculated base_price * multiplier
}

type CreateShowtimeSeatRequest struct {
	ShowtimeID uuid.UUID          `json:"showtime_id"`
	SeatID     uuid.UUID          `json:"seat_id"`
	Status     ShowtimeSeatStatus `json:"status"`
	Price      float64            `json:"price"` // calculated base_price * multiplier
}

type UpdateShowtimeSeatRequest struct {
	ShowtimeID uuid.UUID          `json:"showtime_id"`
	SeatID     uuid.UUID          `json:"seat_id"`
	Status     ShowtimeSeatStatus `json:"status"`
	Price      float64            `json:"price"` // calculated base_price * multiplier
}

func toShowtimeResponse(s *Showtime) ShowtimeResponse {
	if s == nil {
		return ShowtimeResponse{}
	}

	return ShowtimeResponse{
		ID:        s.ID,
		MovieID:   s.MovieID,
		RoomID:    s.RoomID,
		StartTime: s.StartTime,
		EndTime:   s.EndTime,
		BasePrice: s.BasePrice,
		Status:    s.Status,
	}
}

func toShowtimeSeatResponse(s *ShowtimeSeat) ShowtimeSeatResponse {
	if s == nil {
		return ShowtimeSeatResponse{}
	}

	return ShowtimeSeatResponse{
		ID:         s.ID,
		ShowtimeID: s.ShowtimeID,
		SeatID:     s.SeatID,
		Status:     s.Status,
		Price:      s.Price,
	}
}
