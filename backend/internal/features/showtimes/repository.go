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

	// showtime seats
	CreateShowtimeSeat(ctx context.Context, seat *ShowtimeSeat) error
	GetShowtimeSeatByID(ctx context.Context, id uuid.UUID) (*ShowtimeSeat, error)
	GetAllShowtimeSeats(ctx context.Context) ([]*ShowtimeSeat, error)
	UpdateShowtimeSeat(ctx context.Context, seat *ShowtimeSeat) error
	DeleteShowtimeSeat(ctx context.Context, id uuid.UUID) error
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
	query := `
		UPDATE showtimes 
		SET movie_id = $1, room_id = $2, start_time = $3, end_time = $4, base_price = $5, status = $6, updated_at = NOW() 
		WHERE id = $7`

	_, err := r.db.Exec(ctx, query,
		showtime.MovieID,
		showtime.RoomID,
		showtime.StartTime,
		showtime.EndTime,
		showtime.BasePrice,
		showtime.Status,
		showtime.ID,
	)

	return err
}

func (r *repository) DeleteShowtime(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM showtimes WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *repository) CreateShowtimeSeat(ctx context.Context, seat *ShowtimeSeat) error {
	query := `
		INSERT INTO showtime_seats (id, showtime_id, seat_id, status, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())`

	_, err := r.db.Exec(
		ctx,
		query,
		seat.ID,
		seat.ShowtimeID,
		seat.SeatID,
		seat.Status,
		seat.Price,
	)

	return err
}

func (r *repository) GetShowtimeSeatByID(ctx context.Context, id uuid.UUID) (*ShowtimeSeat, error) {
	query := `
		SELECT id, showtime_id, seat_id, status, price, created_at, updated_at 
		FROM showtime_seats 
		WHERE id = $1`

	var s ShowtimeSeat
	err := r.db.QueryRow(ctx, query, id).Scan(
		&s.ID, &s.ShowtimeID, &s.SeatID, &s.Status, &s.Price, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *repository) GetAllShowtimeSeats(ctx context.Context) ([]*ShowtimeSeat, error) {
	query := `SELECT id, showtime_id, seat_id, status, price, created_at, updated_at 
	FROM showtime_seats`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*ShowtimeSeat

	for rows.Next() {
		var s ShowtimeSeat
		err := rows.Scan(
			&s.ID, &s.ShowtimeID, &s.SeatID, &s.Status, &s.Price, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, &s)
	}

	return list, nil
}

func (r *repository) UpdateShowtimeSeat(ctx context.Context, seat *ShowtimeSeat) error {
	query := `
		UPDATE showtime_seats 
		SET showtime_id = $1, seat_id = $2, status = $3, price = $4, updated_at = NOW() 
		WHERE id = $5`

	_, err := r.db.Exec(ctx, query,
		seat.ShowtimeID,
		seat.SeatID,
		seat.Status,
		seat.Price,
		seat.ID,
	)

	return err
}

func (r *repository) DeleteShowtimeSeat(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM showtime_seats WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
