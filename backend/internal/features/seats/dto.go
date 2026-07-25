package seats

import "github.com/google/uuid"

type CreateSeatRequest struct {
	RoomID     uuid.UUID `json:"room_id"`
	SeatTypeID uuid.UUID `json:"seat_type_id"`
	Row        int       `json:"row"`
	Number     int       `json:"number"`
}

type UpdateSeatRequest struct {
	SeatTypeID uuid.UUID `json:"seat_type_id"`
	Row        int       `json:"row"`
	Number     int       `json:"number"`
}

type SeatResponse struct {
	ID         uuid.UUID `json:"id"`
	RoomID     uuid.UUID `json:"room_id"`
	SeatTypeID uuid.UUID `json:"seat_type_id"`
	Row        int       `json:"row"`
	Number     int       `json:"number"`
}

type CreateSeatTypeRequest struct {
	Name       string  `json:"name"`
	Multiplier float64 `json:"multiplier"`
}

type UpdateSeatTypeRequest struct {
	Name       string  `json:"name"`
	Multiplier float64 `json:"multiplier"`
}

type SeatTypeResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Multiplier float64   `json:"multiplier"`
}

func toSeatResponse(seat *Seat) *SeatResponse {
	return &SeatResponse{
		ID:         seat.ID,
		RoomID:     seat.RoomID,
		SeatTypeID: seat.SeatTypeID,
		Row:        seat.Row,
		Number:     seat.Number,
	}
}

func toSeatTypeResponse(seatType *SeatType) *SeatTypeResponse {
	return &SeatTypeResponse{
		ID:         seatType.ID,
		Name:       seatType.Name,
		Multiplier: seatType.Multiplier,
	}
}
