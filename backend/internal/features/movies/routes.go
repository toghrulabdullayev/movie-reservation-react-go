package movies

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	movies := router.Group("/movies", func(c fiber.Ctx) error {
		log.Println("Processing movies request", c.Path(), c.Locals("isAuth"))
		return c.Next()
	})
	movies.Get("/", handler.GetAllMovies)
	movies.Get("/:id", handler.GetMovieByID)
}
