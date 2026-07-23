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

// * Swagger annotations
// @Summary      Get all movies
// @Description  Retrieve a paginated list of movies, sorted by release date descending.
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number (default: 1)"
// @Param        limit  query      int  false  "Number of items per page (default: 10, max: 100)"
// @Success      200    {object}   pagination.PaginatedResponse[MovieResponse]
// @Failure      500    {object}   map[string]string "Internal Server Error"
// @Router       /movies [get]
func (h *Handler) GetAllMovies(c fiber.Ctx) error {
	// extract and sanitize page/limit params (defaults to 10 per page, max 100)
	params := pagination.GetPaginationParams(c, 10, 100)

	// fetch movies and total records
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

	// wrap with paginated structure
	paginatedResponse := pagination.NewPaginatedResponse(movieDTOs, totalRecords, params)

	return c.JSON(paginatedResponse)
}

// * Swagger annotations
// GetMovieByID retrieves a single movie by its UUID.
// @Summary      Get movie by ID
// @Description  Retrieve details of a single movie including its associated genres.
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Movie UUID" format(uuid)
// @Success      200  {object}  MovieResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid UUID format"
// @Failure      404  {object}  map[string]string "Not Found - Movie not found"
// @Router       /movies/{id} [get]
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
