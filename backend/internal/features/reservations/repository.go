package reservations

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/pagination"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	CreateReservation(ctx context.Context, reservation *Reservation) error
	GetReservationByID(ctx context.Context, id uuid.UUID) (*Reservation, error)
	GetAllReservations(ctx context.Context, params pagination.PaginationParams) ([]*Reservation, int, error)
	UpdateReservation(ctx context.Context, reservation *Reservation) error
	DeleteReservation(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateReservation(ctx context.Context, reservation *Reservation) error {
	query := `
		INSERT INTO reservations (id, user_id, showtime_id, status, total_price, reserved_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		reservation.ID,
		reservation.UserID,
		reservation.ShowtimeID,
		reservation.Status,
		reservation.TotalPrice,
	)

	return err
}

func (r *repository) GetReservationByID(ctx context.Context, id uuid.UUID) (*Reservation, error) {
	query := `
		SELECT id, user_id, showtime_id, status, total_price, reserved_at, created_at, updated_at 
		FROM reservations 
		WHERE id = $1`

	var reservation Reservation
	err := r.db.QueryRow(ctx, query, id).Scan(
		&reservation.ID,
		&reservation.UserID,
		&reservation.ShowtimeID,
		&reservation.Status,
		&reservation.TotalPrice,
		&reservation.ReservedAt,
		&reservation.CreatedAt,
		&reservation.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *repository) GetAllReservations(ctx context.Context, params pagination.PaginationParams) ([]*Reservation, int, error) {
	totalRecordsQuery := `SELECT COUNT(*) FROM reservations`
	query := `
		SELECT id, user_id, showtime_id, status, total_price, reserved_at, created_at, updated_at 
		FROM reservations 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	totalRecords := 0
	err := r.db.QueryRow(ctx, totalRecordsQuery).Scan(&totalRecords)
	if err != nil {
		return nil, 0, err
	}

	if totalRecords == 0 {
		return []*Reservation{}, 0, nil
	}

	rows, err := r.db.Query(ctx, query, params.Limit, params.Offset())
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var reservations []*Reservation
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(
			&reservation.ID,
			&reservation.UserID,
			&reservation.ShowtimeID,
			&reservation.Status,
			&reservation.TotalPrice,
			&reservation.ReservedAt,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		reservations = append(reservations, &reservation)
	}

	return reservations, totalRecords, nil
}

func (r *repository) UpdateReservation(ctx context.Context, reservation *Reservation) error {
	query := `
		UPDATE reservations 
		SET user_id = $1, showtime_id = $2, status = $3, total_price = $4, updated_at = NOW() 
		WHERE id = $5`

	_, err := r.db.Exec(
		ctx,
		query,
		reservation.UserID,
		reservation.ShowtimeID,
		reservation.Status,
		reservation.TotalPrice,
		reservation.ID,
	)

	return err
}

func (r *repository) DeleteReservation(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM reservations WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
