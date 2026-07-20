package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/showtimes"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
)

type Dependencies struct {
	MovieHandler    *movies.Handler
	ShowtimeHandler *showtimes.Handler
}

func BuildDependencies(cfg *config.Config, pgPool *pgxpool.Pool, rdb *redis.Client) *Dependencies {
	movieRepository := movies.NewRepository(pgPool)
	movieService := movies.NewService(movieRepository)
	movieHandler := movies.NewHandler(movieService)

	showtimeRepository := showtimes.NewRepository(pgPool)
	showtimeService := showtimes.NewService(showtimeRepository)
	showtimeHandler := showtimes.NewHandler(showtimeService)

	return &Dependencies{
		MovieHandler:    movieHandler,
		ShowtimeHandler: showtimeHandler,
	}
}
