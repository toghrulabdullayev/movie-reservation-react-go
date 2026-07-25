package seats

import (
	"context"

	"github.com/google/uuid"
)

var _ Service = (*service)(nil)

type Service interface {
	// seats
	CreateSeat(ctx context.Context, seat *Seat) error
	GetSeatByID(ctx context.Context, id uuid.UUID) (*Seat, error)
	GetAllSeats(ctx context.Context) ([]*Seat, error)
	UpdateSeat(ctx context.Context, seat *Seat) error
	DeleteSeat(ctx context.Context, id uuid.UUID) error

	// seat types
	CreateSeatType(ctx context.Context, seatType *SeatType) error
	GetSeatTypeByID(ctx context.Context, id uuid.UUID) (*SeatType, error)
	GetAllSeatTypes(ctx context.Context) ([]*SeatType, error)
	UpdateSeatType(ctx context.Context, seatType *SeatType) error
	DeleteSeatType(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// seats
func (s *service) CreateSeat(ctx context.Context, seat *Seat) error {
	return s.repo.CreateSeat(ctx, seat)
}

func (s *service) GetSeatByID(ctx context.Context, id uuid.UUID) (*Seat, error) {
	return s.repo.GetSeatByID(ctx, id)
}

func (s *service) GetAllSeats(ctx context.Context) ([]*Seat, error) {
	return s.repo.GetAllSeats(ctx)
}

func (s *service) UpdateSeat(ctx context.Context, seat *Seat) error {
	return s.repo.UpdateSeat(ctx, seat)
}

func (s *service) DeleteSeat(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSeat(ctx, id)
}

// seat types
func (s *service) CreateSeatType(ctx context.Context, seatType *SeatType) error {
	return s.repo.CreateSeatType(ctx, seatType)
}

func (s *service) GetSeatTypeByID(ctx context.Context, id uuid.UUID) (*SeatType, error) {
	return s.repo.GetSeatTypeByID(ctx, id)
}

func (s *service) GetAllSeatTypes(ctx context.Context) ([]*SeatType, error) {
	return s.repo.GetAllSeatTypes(ctx)
}

func (s *service) UpdateSeatType(ctx context.Context, seatType *SeatType) error {
	return s.repo.UpdateSeatType(ctx, seatType)
}

func (s *service) DeleteSeatType(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSeatType(ctx, id)
}
