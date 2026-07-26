package server

import (
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	_ "github.com/toghrulabdullayev/movie-reservation-react-go/backend/docs"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/auth"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/reservations"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/rooms"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/seats"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/showtimes"
	// "github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/middleware"
)

func registerRoutes(app *fiber.App, deps *Dependencies) {
	// health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// swagger UI
	app.Get("/swagger/*", swaggo.HandlerDefault)

	// // customized swagger UI
	// app.Get("/swagger/*", swaggo.New(swaggo.Config{
	// 	Title: "Movie Reservation API",
	// }))

	// API v1 routes
	v1 := app.Group("/api/v1")
	// protected := v1.Group("", middleware.Auth()) //! auth protects all /api/v1 routes

	// todo: deps.MovieHandler.RegisterRoutes(v1)
	auth.RegisterRoutes(v1, deps.AuthHandler)
	movies.RegisterRoutes(v1, deps.MovieHandler)
	reservations.RegisterRoutes(v1, deps.ReservationHandler)
	rooms.RegisterRoutes(v1, deps.RoomHandler)
	seats.RegisterRoutes(v1, deps.SeatHandler)
	showtimes.RegisterRoutes(v1, deps.ShowtimeHandler)

	// deps.ReservationHandler.RegisterRoutes(protected)

	// admin := v1.Group("", middleware.Auth(), middleware.RBAC("admin"))
	// deps.ReportHandler.RegisterRoutes(admin)
	// deps.ShowtimeHandler.RegisterRoutes(admin) // if showtime mgmt is admin-only
}
