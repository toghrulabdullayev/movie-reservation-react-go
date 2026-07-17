package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/features/movies"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/platform/middleware"
)

func registerRoutes(app *fiber.App, deps *Dependencies) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	v1 := app.Group("/api/v1")
	protected := v1.Group("", middleware.Auth())

	// todo: deps.MovieHandler.RegisterRoutes(v1)
	movies.RegisterRoutes(protected, deps.MovieHandler)

	// deps.ReservationHandler.RegisterRoutes(protected)

	// admin := v1.Group("", middleware.Auth(), middleware.RBAC("admin"))
	// deps.ReportHandler.RegisterRoutes(admin)
	// deps.ShowtimeHandler.RegisterRoutes(admin) // if showtime mgmt is admin-only
}
