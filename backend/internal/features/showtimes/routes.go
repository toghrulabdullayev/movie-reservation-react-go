package showtimes

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	showtimes := router.Group("/showtimes", func(c fiber.Ctx) error {
		log.Println("Processing showtimes request", c.Path(), c.Locals("isAuth"))
		return c.Next()
	})

	// showtime seats
	showtimes.Route("/seats", func(r fiber.Router) {
		r.Get("/", handler.GetAllShowtimeSeats)
		r.Get("/:id", handler.GetShowtimeSeatByID)
		r.Post("/", handler.CreateShowtimeSeat)
		r.Put("/:id", handler.UpdateShowtimeSeat)
		r.Delete("/:id", handler.DeleteShowtimeSeat)
	})

	// showtime
	showtimes.Get("/", handler.GetAllShowtimes)
	showtimes.Get("/:id", handler.GetShowtimeByID)
	showtimes.Post("/", handler.CreateShowtime)
	showtimes.Put("/:id", handler.UpdateShowtime)
	showtimes.Delete("/:id", handler.DeleteShowtime)
}
