package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/auth"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/reservations"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/rooms"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/seats"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/showtimes"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/config"
)

type Dependencies struct {
	Config *config.Config

	AuthHandler        *auth.Handler
	MovieHandler       *movies.Handler
	ReservationHandler *reservations.Handler
	RoomHandler        *rooms.Handler
	SeatHandler        *seats.Handler
	ShowtimeHandler    *showtimes.Handler
}

func BuildDependencies(cfg *config.Config, pgPool *pgxpool.Pool, rdb *redis.Client) *Dependencies {
	authRepository := auth.NewRepository(pgPool)
	authService := auth.NewService(authRepository, rdb, cfg)
	authHandler := auth.NewHandler(authService, cfg)

	movieRepository := movies.NewRepository(pgPool)
	movieService := movies.NewService(movieRepository)
	movieHandler := movies.NewHandler(movieService)

	reservationRepository := reservations.NewRepository(pgPool)
	reservationService := reservations.NewService(reservationRepository)
	reservationHandler := reservations.NewHandler(reservationService)

	roomRepository := rooms.NewRepository(pgPool)
	roomService := rooms.NewService(roomRepository)
	roomHandler := rooms.NewHandler(roomService)

	seatRepository := seats.NewRepository(pgPool)
	seatService := seats.NewService(seatRepository)
	seatHandler := seats.NewHandler(seatService)

	showtimeRepository := showtimes.NewRepository(pgPool)
	showtimeService := showtimes.NewService(showtimeRepository)
	showtimeHandler := showtimes.NewHandler(showtimeService)

	return &Dependencies{
		Config: cfg,

		AuthHandler:        authHandler,
		MovieHandler:       movieHandler,
		ReservationHandler: reservationHandler,
		RoomHandler:        roomHandler,
		SeatHandler:        seatHandler,
		ShowtimeHandler:    showtimeHandler,
	}
}
