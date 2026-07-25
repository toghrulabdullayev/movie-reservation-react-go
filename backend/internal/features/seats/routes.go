package seats

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	seats := router.Group("/seats", func(c fiber.Ctx) error {
		log.Println("Processing seats request", c.Path(), c.Locals("isAuth"))
		return c.Next()
	})

	//! order of routes matters like in nodejs + express backends
	// seat types routes
	seats.Route("/types", func(r fiber.Router) {
		r.Get("/", handler.GetAllSeatTypes)
		r.Get("/:id", handler.GetSeatTypeByID)
		r.Post("/", handler.CreateSeatType)
		r.Put("/:id", handler.UpdateSeatType)
		r.Delete("/:id", handler.DeleteSeatType)
	})

	// seats routes
	seats.Get("/", handler.GetAllSeats)
	seats.Get("/:id", handler.GetSeatByID)
	seats.Post("/", handler.CreateSeat)
	seats.Put("/:id", handler.UpdateSeat)
	seats.Delete("/:id", handler.DeleteSeat)
}
