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

// add to dto.go
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
