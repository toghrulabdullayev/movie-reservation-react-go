package movies

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	movies := router.Group("/movies")
	movies.Get("/", handler.GetMovies)
	movies.Get("/:id", handler.GetMovieByID)
	movies.Post("/", handler.CreateMovie)
	movies.Put("/:id", handler.UpdateMovie)
	movies.Delete("/:id", handler.DeleteMovie)
}
