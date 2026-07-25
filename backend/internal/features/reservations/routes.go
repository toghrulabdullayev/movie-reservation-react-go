package reservations

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	reservations := router.Group("/reservations", func(c fiber.Ctx) error {
		log.Println("Processing reservations request", c.Path(), c.Locals("isAuth"))
		return c.Next()
	})
	reservations.Get("/", handler.GetAllReservations)
	reservations.Get("/:id", handler.GetReservationByID)
	reservations.Post("/", handler.CreateReservation)
	reservations.Put("/:id", handler.UpdateReservation)
	reservations.Delete("/:id", handler.DeleteReservation)
}

