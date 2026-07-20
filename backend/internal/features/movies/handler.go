package movies

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAllMovies(c fiber.Ctx) error {
	movies, err := h.service.GetAllMovies(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve movies",
		})
	}

	var response []*MovieResponse
	for _, m := range movies {
		response = append(response, toMovieResponse(m))
	}

	return c.JSON(response)
}

func (h *Handler) GetMovieByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid movie ID format",
		})
	}

	movie, err := h.service.GetMovieByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "movie not found",
		})
	}

	return c.JSON(toMovieResponse(movie))
}
