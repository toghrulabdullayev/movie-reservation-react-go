package pagination

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// PaginationParams holds standard pagination criteria parsed from the request
type PaginationParams struct {
	Page  int
	Limit int
}

// Offset calculates the SQL OFFSET value
func (p PaginationParams) Offset() int {
	return (p.Page - 1) * p.Limit
}

// GetPaginationParams extracts and sanitizes page and limit from query parameters
func GetPaginationParams(c fiber.Ctx, defaultLimit, maxLimit int) PaginationParams {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	limit := defaultLimit
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
		if limit > maxLimit {
			limit = maxLimit
		}
	}

	return PaginationParams{
		Page:  page,
		Limit: limit,
	}
}

// PaginationMetadata holds standard metadata about a paginated request
type PaginationMetadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
}

// PaginatedResponse wraps generic slices of data with pagination metadata
type PaginatedResponse[T any] struct {
	Metadata PaginationMetadata `json:"metadata"`
	Data     []T                `json:"data"`
}

// NewPaginatedResponse constructs a standard paginated response payload
func NewPaginatedResponse[T any](data []T, totalRecords int, params PaginationParams) PaginatedResponse[T] {
	totalPages := 0
	if totalRecords > 0 {
		totalPages = (totalRecords + params.Limit - 1) / params.Limit
	}

	return PaginatedResponse[T]{
		Metadata: PaginationMetadata{
			CurrentPage:  params.Page,
			PageSize:     params.Limit,
			TotalRecords: totalRecords,
			TotalPages:   totalPages,
		},
		Data: data,
	}
}
