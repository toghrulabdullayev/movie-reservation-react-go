package showtimes

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// compile-time check to ensure repository implements Repository.
var _ Repository = (*repository)(nil)

type Repository interface {
	// showtimes
	CreateShowtime(ctx context.Context, showtime *Showtime) error
	GetShowtimeByID(ctx context.Context, id uuid.UUID) (*Showtime, error)
	GetAllShowtimes(ctx context.Context) ([]*Showtime, error)
	UpdateShowtime(ctx context.Context, showtime *Showtime) error
	DeleteShowtime(ctx context.Context, id uuid.UUID) error
}

// concrete implementation — lowercase, unexported
type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

// showtimes
func (r *repository) CreateShowtime(ctx context.Context, showtime *Showtime) error {
	query := `
		INSERT INTO showtimes (id, movie_id, room_id, start_time, end_time, base_price, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		showtime.ID,
		showtime.MovieID,
		showtime.RoomID,
		showtime.StartTime,
		showtime.EndTime,
		showtime.BasePrice,
		showtime.Status,
	)

	return err
}

func (r *repository) GetShowtimeByID(ctx context.Context, id uuid.UUID) (*Showtime, error) {
	query := `
		SELECT id, movie_id, room_id, start_time, end_time, base_price, status, created_at, updated_at 
		FROM showtimes 
		WHERE id = $1`

	var s Showtime
	err := r.db.QueryRow(ctx, query, id).Scan(
		&s.ID, &s.MovieID, &s.RoomID, &s.StartTime,
		&s.EndTime, &s.BasePrice, &s.Status, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *repository) GetAllShowtimes(ctx context.Context) ([]*Showtime, error) {
	query := `SELECT id, movie_id, room_id, start_time, end_time, base_price, status, created_at, updated_at 
	FROM showtimes`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*Showtime

	for rows.Next() {
		var s Showtime
		err := rows.Scan(
			&s.ID, &s.MovieID, &s.RoomID, &s.StartTime,
			&s.EndTime, &s.BasePrice, &s.Status, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, &s)
	}

	return list, nil
}

func (r *repository) UpdateShowtime(ctx context.Context, showtime *Showtime) error {
	panic("not implemented")
}

func (r *repository) DeleteShowtime(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}
