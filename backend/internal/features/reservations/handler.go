package reservations

import (
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
// @Summary      Create a new reservation
// @Description  Create a new reservation with the provided details.
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        reservation  body      CreateReservationRequest  true  "Reservation details"
// @Success      201          {object}  ReservationResponse
// @Failure      400          {object}  map[string]string "Bad Request - Invalid input"
// @Failure      500          {object}  map[string]string "Internal Server Error - Failed to create reservation"
// @Router       /reservations [post]
func (h *Handler) CreateReservation(c fiber.Ctx) error {
	var req CreateReservationRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	reservation, err := h.service.CreateReservation(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create reservation",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(toReservationResponse(reservation))
}

// * Swagger annotations
// @Summary      Get reservation by ID
// @Description  Retrieve the details of a single reservation by its UUID.
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Reservation UUID" format(uuid)
// @Success      200  {object}  ReservationResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Reservation not found"
// @Router       /reservations/{id} [get]
func (h *Handler) GetReservationByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid reservation ID format",
		})
	}

	reservation, err := h.service.GetReservationByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve reservation",
		})
	}

	return c.JSON(toReservationResponse(reservation))
}

// * Swagger annotations
// @Summary      Get all reservations
// @Description  Retrieve a list of all reservations with pagination.
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        page    query     int  false  "Page number" mininum(1)
// @Param        limit   query     int  false  "Number of items per page" mininum(1) maximum(100)
// @Success      200     {object}  pagination.PaginatedResponse[ReservationResponse]
// @Failure      500     {object}  map[string]string "Internal Server Error - Failed to retrieve reservations"
// @Router       /reservations [get]
func (h *Handler) GetAllReservations(c fiber.Ctx) error {
	params := pagination.GetPaginationParams(c, 10, 100)

	reservations, totalRecords, err := h.service.GetAllReservations(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve reservations",
		})
	}

	reservationResponses := make([]ReservationResponse, len(reservations))
	for _, reservation := range reservations {
		reservationResponses = append(reservationResponses, toReservationResponse(reservation))
	}

	paginatedResponse := pagination.NewPaginatedResponse(reservationResponses, totalRecords, params)

	return c.JSON(paginatedResponse)
}

// * Swagger annotations
// @Summary      Update a reservation
// @Description  Update the details of an existing reservation.
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Reservation UUID" format(uuid)
// @Param        reservation  body      UpdateReservationRequest  true  "Updated reservation details"
// @Success      200  {object}  ReservationResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid input"
// @Failure      404  {object}  map[string]string "Not Found - Reservation not found"
// @Router       /reservations/{id} [put]
func (h *Handler) UpdateReservation(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid reservation ID format",
		})
	}

	var req UpdateReservationRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	reservation, err := h.service.UpdateReservation(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update reservation",
		})
	}

	return c.JSON(toReservationResponse(reservation))
}

// * Swagger annotations
// @Summary      Delete a reservation
// @Description  Delete an existing reservation by its UUID.
// @Tags         reservations
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Reservation UUID" format(uuid)
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Reservation not found"
// @Router       /reservations/{id} [delete]
func (h *Handler) DeleteReservation(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid reservation ID format",
		})
	}

	if err := h.service.DeleteReservation(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete reservation",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
