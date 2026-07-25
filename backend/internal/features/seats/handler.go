package seats

import (
	"log"

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

// seats
// * Swagger annotations
// @Summary      Create a new seat
// @Description  Create a new seat with the provided details.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param        seat  body      CreateSeatRequest  true  "Seat details"
// @Success      201   {object}  SeatResponse
// @Failure      400   {object}  map[string]string "Bad Request - Invalid input"
// @Failure      500   {object}  map[string]string "Internal Server Error - Failed to create seat"
// @Router       /seats [post]
func (h *Handler) CreateSeat(c fiber.Ctx) error {
	var req CreateSeatRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	seat := &Seat{
		ID:         uuid.New(),
		RoomID:     req.RoomID,
		SeatTypeID: req.SeatTypeID,
		Row:        req.Row,
		Number:     req.Number,
	}

	if err := h.service.CreateSeat(c.Context(), seat); err != nil {
		log.Printf("Failed to create seat: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create seat",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(toSeatResponse(seat))
}

// * Swagger annotations
// @Summary      Get seat by ID
// @Description  Retrieve the details of a single seat by its UUID.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Seat UUID" format(uuid)
// @Success      200  {object}  SeatResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Seat not found"
// @Router       /seats/{id} [get]
func (h *Handler) GetSeatByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid seat ID format",
		})
	}

	seat, err := h.service.GetSeatByID(c.Context(), id)
	if err != nil {
		log.Printf("Failed to get seat by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get seat",
		})
	}

	if seat == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Seat not found",
		})
	}

	return c.JSON(toSeatResponse(seat))
}

// * Swagger annotations
// @Summary      Get all seats
// @Description  Retrieve a list of all seats.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Success      200  {array}   SeatResponse
// @Failure      500  {object}  map[string]string "Internal Server Error - Failed to get seats"
// @Router       /seats [get]
func (h *Handler) GetAllSeats(c fiber.Ctx) error {
	seats, err := h.service.GetAllSeats(c.Context())
	if err != nil {
		log.Printf("Failed to get all seats: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get seats",
		})
	}

	var seatResponses []*SeatResponse
	for _, seat := range seats {
		seatResponses = append(seatResponses, toSeatResponse(seat))
	}

	return c.JSON(seatResponses)
}

// * Swagger annotations
// @Summary      Update seat
// @Description  Update the details of an existing seat.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Seat UUID" format(uuid)
// @Param        seat body      CreateSeatRequest  true  "Updated seat details"
// @Success      200  {object}  SeatResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid input"
// @Failure      404  {object}  map[string]string "Not Found - Seat not found"
// @Router       /seats/{id} [put]
func (h *Handler) UpdateSeat(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid seat ID format",
		})
	}

	var req CreateSeatRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	seat := &Seat{
		ID:         id,
		RoomID:     req.RoomID,
		SeatTypeID: req.SeatTypeID,
		Row:        req.Row,
		Number:     req.Number,
	}

	if err := h.service.UpdateSeat(c.Context(), seat); err != nil {
		log.Printf("Failed to update seat: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update seat",
		})
	}

	return c.JSON(toSeatResponse(seat))
}

// * Swagger annotations
// @Summary      Delete seat
// @Description  Delete a seat by its UUID.
// @Tags         seats
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Seat UUID" format(uuid)
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Seat not found"
// @Router       /seats/{id} [delete]
func (h *Handler) DeleteSeat(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid seat ID format",
		})
	}

	if err := h.service.DeleteSeat(c.Context(), id); err != nil {
		log.Printf("Failed to delete seat: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete seat",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// seat types
// * Swagger annotations
// @Summary      Get seat type by ID
// @Description  Retrieve the details of a single seat type by its UUID.
// @Tags         seat-types
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Seat Type UUID" format(uuid)
// @Success      200  {object}  SeatTypeResponse
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Seat type not found"
// @Router       /seats/types/{id} [get]
func (h *Handler) GetSeatTypeByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid seat type ID format",
		})
	}

	seatType, err := h.service.GetSeatTypeByID(c.Context(), id)
	if err != nil {
		log.Printf("Failed to get seat type by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get seat type",
		})
	}

	if seatType == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Seat type not found",
		})
	}

	return c.JSON(seatType)
}

// * Swagger annotations
// @Summary      Get all seat types
// @Description  Retrieve a list of all seat types.
// @Tags         seat-types
// @Accept       json
// @Produce      json
// @Success      200  {array}   SeatTypeResponse
// @Failure      500  {object}  map[string]string "Internal Server Error - Failed to get seat types"
// @Router       /seats/types [get]
func (h *Handler) GetAllSeatTypes(c fiber.Ctx) error {
	seatTypes, err := h.service.GetAllSeatTypes(c.Context())
	if err != nil {
		log.Printf("Failed to get all seat types: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get seat types",
		})
	}

	return c.JSON(seatTypes)
}

// * Swagger annotations
// @Summary      Create a new seat type
// @Description  Create a new seat type with the provided details.
// @Tags         seat-types
// @Accept       json
// @Produce      json
// @Param        seatType  body      CreateSeatTypeRequest  true  "Seat type details"
// @Success      201       {object}  SeatTypeResponse
// @Failure      400       {object}  map[string]string "Bad Request - Invalid input"
// @Failure      500       {object}  map[string]string "Internal Server Error - Failed to create seat type"
// @Router       /seats/types [post]
func (h *Handler) CreateSeatType(c fiber.Ctx) error {
	var seatType SeatType
	if err := c.Bind().Body(&seatType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	seatType.ID = uuid.New()

	if err := h.service.CreateSeatType(c.Context(), &seatType); err != nil {
		log.Printf("Failed to create seat type: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create seat type",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(seatType)
}

// * Swagger annotations
// @Summary      Update seat type
// @Description  Update the details of an existing seat type.
// @Tags         seat-types
// @Accept       json
// @Produce      json
// @Param        id        path      string    true  "Seat Type UUID" format(uuid)
// @Param        seatType  body      UpdateSeatTypeRequest  true  "Updated seat type details"
// @Success      200       {object}  SeatTypeResponse
// @Failure      400       {object}  map[string]string "Bad Request - Invalid input"
// @Failure      404       {object}  map[string]string "Not Found - Seat type not found"
// @Router       /seats/types/{id} [put]
func (h *Handler) UpdateSeatType(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid seat type ID format",
		})
	}

	var seatType SeatType
	if err := c.Bind().Body(&seatType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	seatType.ID = id

	if err := h.service.UpdateSeatType(c.Context(), &seatType); err != nil {
		log.Printf("Failed to update seat type: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update seat type",
		})
	}

	return c.JSON(seatType)
}

// * Swagger annotations
// @Summary      Delete seat type
// @Description  Delete a seat type by its UUID.
// @Tags         seat-types
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Seat Type UUID" format(uuid)
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string "Bad Request - Invalid ID format"
// @Failure      404  {object}  map[string]string "Not Found - Seat type not found"
// @Router       /seats/types/{id} [delete]
func (h *Handler) DeleteSeatType(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid seat type ID format",
		})
	}

	if err := h.service.DeleteSeatType(c.Context(), id); err != nil {
		log.Printf("Failed to delete seat type: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete seat type",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
