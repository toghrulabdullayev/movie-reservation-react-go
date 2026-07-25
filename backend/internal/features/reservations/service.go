package reservations

import (
	"context"

	"github.com/google/uuid"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/pagination"
)

var _ Service = (*service)(nil)

type Service interface {
	// reservations
	CreateReservation(ctx context.Context, req CreateReservationRequest) (*Reservation, error)
	GetReservationByID(ctx context.Context, id uuid.UUID) (*Reservation, error)
	GetAllReservations(ctx context.Context, params pagination.PaginationParams) ([]*Reservation, int, error)
	UpdateReservation(ctx context.Context, id uuid.UUID, req UpdateReservationRequest) (*Reservation, error)
	DeleteReservation(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// reservations
func (s *service) CreateReservation(ctx context.Context, req CreateReservationRequest) (*Reservation, error) {
	reservation := &Reservation{
		ID:         uuid.New(),
		UserID:     req.UserID,
		ShowtimeID: req.ShowtimeID,
		Status:     ReservationPending,
		TotalPrice: req.TotalPrice,
	}

	if err := s.repo.CreateReservation(ctx, reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *service) GetReservationByID(ctx context.Context, id uuid.UUID) (*Reservation, error) {
	return s.repo.GetReservationByID(ctx, id)
}

func (s *service) GetAllReservations(ctx context.Context, params pagination.PaginationParams) ([]*Reservation, int, error) {
	return s.repo.GetAllReservations(ctx, params)
}

func (s *service) UpdateReservation(ctx context.Context, id uuid.UUID, req UpdateReservationRequest) (*Reservation, error) {
	reservation := &Reservation{
		ID:         id,
		UserID:     req.UserID,
		ShowtimeID: req.ShowtimeID,
		Status:     req.Status,
		TotalPrice: req.TotalPrice,
	}

	if err := s.repo.UpdateReservation(ctx, reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *service) DeleteReservation(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteReservation(ctx, id)
}
