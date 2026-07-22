package movies

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/toghrulabdullayev/movie-reservation-react-go/backend/internal/shared/pagination"
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
	// extract and sanitize page/limit params (defaults to 10 per page, max 100)
	params := pagination.GetPaginationParams(c, 10, 100)

	// fetch movies and total records from the service
	moviesList, totalRecords, err := h.service.GetAllMovies(c.Context(), params)
	if err != nil {
		log.Printf("[API ERROR] GetAllMovies failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve movies",
		})
	}

	// map entities to response DTOs
	movieDTOs := make([]*MovieResponse, 0, len(moviesList))
	for _, m := range moviesList {
		movieDTOs = append(movieDTOs, toMovieResponse(m))
	}

	// wrap with our generic paginated structure
	paginatedResponse := pagination.NewPaginatedResponse(movieDTOs, totalRecords, params)

	return c.JSON(paginatedResponse)
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
		log.Printf("[API ERROR] GetMovieByID failed: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "movie not found",
		})
	}

	return c.JSON(toMovieResponse(movie))
}
