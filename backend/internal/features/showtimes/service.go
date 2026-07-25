package showtimes

import (
	"context"

	"github.com/google/uuid"
)

var _ Service = (*service)(nil)

type Service interface {
	// showtimes
	CreateShowtime(ctx context.Context, req CreateShowtimeRequest) (*Showtime, error)
	GetShowtimeByID(ctx context.Context, id uuid.UUID) (*Showtime, error)
	GetAllShowtimes(ctx context.Context) ([]*Showtime, error)
	UpdateShowtime(ctx context.Context, id uuid.UUID, req UpdateShowtimeRequest) error
	DeleteShowtime(ctx context.Context, id uuid.UUID) error

	// showtime seats
	CreateShowtimeSeat(ctx context.Context, req CreateShowtimeSeatRequest) error
	GetShowtimeSeatByID(ctx context.Context, id uuid.UUID) (*ShowtimeSeat, error)
	GetAllShowtimeSeats(ctx context.Context) ([]*ShowtimeSeat, error)
	UpdateShowtimeSeat(ctx context.Context, id uuid.UUID, req UpdateShowtimeSeatRequest) error
	DeleteShowtimeSeat(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// showtimes
func (s *service) CreateShowtime(ctx context.Context, req CreateShowtimeRequest) (*Showtime, error) {
	showtime := &Showtime{
		ID:        uuid.New(),
		MovieID:   req.MovieID,
		RoomID:    req.RoomID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		BasePrice: req.BasePrice,
		Status:    req.Status,
	}

	if err := s.repo.CreateShowtime(ctx, showtime); err != nil {
		return nil, err
	}

	return showtime, nil
}

func (s *service) GetShowtimeByID(ctx context.Context, id uuid.UUID) (*Showtime, error) {
	return s.repo.GetShowtimeByID(ctx, id)
}

func (s *service) GetAllShowtimes(ctx context.Context) ([]*Showtime, error) {
	return s.repo.GetAllShowtimes(ctx)
}

func (s *service) UpdateShowtime(ctx context.Context, id uuid.UUID, req UpdateShowtimeRequest) error {
	showtime := &Showtime{
		ID:        id,
		MovieID:   req.MovieID,
		RoomID:    req.RoomID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		BasePrice: req.BasePrice,
		Status:    req.Status,
	}

	return s.repo.UpdateShowtime(ctx, showtime)
}

func (s *service) DeleteShowtime(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteShowtime(ctx, id)
}

func (s *service) CreateShowtimeSeat(ctx context.Context, req CreateShowtimeSeatRequest) error {
	seat := &ShowtimeSeat{
		ID:         uuid.New(),
		ShowtimeID: req.ShowtimeID,
		SeatID:     req.SeatID,
		Status:     req.Status,
		Price:      req.Price,
	}

	return s.repo.CreateShowtimeSeat(ctx, seat)
}

func (s *service) GetShowtimeSeatByID(ctx context.Context, id uuid.UUID) (*ShowtimeSeat, error) {
	return s.repo.GetShowtimeSeatByID(ctx, id)
}

func (s *service) GetAllShowtimeSeats(ctx context.Context) ([]*ShowtimeSeat, error) {
	return s.repo.GetAllShowtimeSeats(ctx)
}

func (s *service) UpdateShowtimeSeat(ctx context.Context, id uuid.UUID, req UpdateShowtimeSeatRequest) error {
	seat := &ShowtimeSeat{
		ID:         id,
		ShowtimeID: req.ShowtimeID,
		SeatID:     req.SeatID,
		Status:     req.Status,
		Price:      req.Price,
	}

	return s.repo.UpdateShowtimeSeat(ctx, seat)
}

func (s *service) DeleteShowtimeSeat(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteShowtimeSeat(ctx, id)
}
