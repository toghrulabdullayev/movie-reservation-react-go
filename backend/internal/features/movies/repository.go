package movies

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error)
	GetAllMovies(ctx context.Context) ([]*Movie, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetMovieByID(ctx context.Context, id uuid.UUID) (*Movie, error) {
	query := `
		SELECT m.id, m.tmdb_id, m.title, 
		       COALESCE(m.overview, '') as overview, 
		       COALESCE(m.poster_path, '') as poster_path, 
		       COALESCE(m.backdrop_path, '') as backdrop_path, 
		       COALESCE(m.release_date, '0001-01-01'::date) as release_date, 
		       COALESCE(m.runtime, 0) as runtime, 
		       COALESCE(m.vote_average, 0.0) as vote_average, 
		       m.created_at, m.updated_at,
		       COALESCE(array_to_json(array_agg(g.name) FILTER (WHERE g.name IS NOT NULL)), '[]') as genres
		FROM movies m
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		WHERE m.id = $1
		GROUP BY m.id`

	var m Movie
	var genresJSON []byte

	err := r.db.QueryRow(ctx, query, id).Scan(
		&m.ID, &m.TMDBID, &m.Title, &m.Overview, &m.PosterPath, &m.BackdropPath,
		&m.ReleaseDate, &m.Runtime, &m.VoteAverage, &m.CreatedAt, &m.UpdatedAt,
		&genresJSON,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	_ = json.Unmarshal(genresJSON, &m.Genres)
	return &m, nil
}

func (r *repository) GetAllMovies(ctx context.Context) ([]*Movie, error) {
	query := `
		SELECT m.id, m.tmdb_id, m.title, 
		       COALESCE(m.overview, '') as overview, 
		       COALESCE(m.poster_path, '') as poster_path, 
		       COALESCE(m.backdrop_path, '') as backdrop_path, 
		       COALESCE(m.release_date, '0001-01-01'::date) as release_date, 
		       COALESCE(m.runtime, 0) as runtime, 
		       COALESCE(m.vote_average, 0.0) as vote_average, 
		       m.created_at, m.updated_at,
		       COALESCE(array_to_json(array_agg(g.name) FILTER (WHERE g.name IS NOT NULL)), '[]') as genres
		FROM movies m
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		GROUP BY m.id
		ORDER BY m.release_date DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*Movie
	for rows.Next() {
		var m Movie
		var genresJSON []byte
		err := rows.Scan(
			&m.ID, &m.TMDBID, &m.Title, &m.Overview, &m.PosterPath, &m.BackdropPath,
			&m.ReleaseDate, &m.Runtime, &m.VoteAverage, &m.CreatedAt, &m.UpdatedAt,
			&genresJSON,
		)
		if err != nil {
			return nil, err
		}
		_ = json.Unmarshal(genresJSON, &m.Genres)
		list = append(list, &m)
	}

	return list, nil
}
