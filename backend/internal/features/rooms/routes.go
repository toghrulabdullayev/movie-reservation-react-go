package rooms

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	rooms := router.Group("/rooms", func(c fiber.Ctx) error {
		log.Println("Processing rooms request", c.Path(), c.Locals("isAuth"))
		return c.Next()
	})
	rooms.Get("/", handler.GetAllRooms)
	rooms.Get("/:id", handler.GetRoomByID)
	rooms.Post("/", handler.CreateRoom)
	rooms.Put("/:id", handler.UpdateRoom)
	rooms.Delete("/:id", handler.DeleteRoom)
}
