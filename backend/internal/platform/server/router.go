package server

import (
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	_ "github.com/toghrulabdullayev/movie-reservation-react-go/backend/docs"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
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
	movies.RegisterRoutes(v1, deps.MovieHandler)
	showtimes.RegisterRoutes(v1, deps.ShowtimeHandler)

	// deps.ReservationHandler.RegisterRoutes(protected)

	// admin := v1.Group("", middleware.Auth(), middleware.RBAC("admin"))
	// deps.ReportHandler.RegisterRoutes(admin)
	// deps.ShowtimeHandler.RegisterRoutes(admin) // if showtime mgmt is admin-only
}
