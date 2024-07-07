package response

import (
	"math"
)

// Paginator holds the pagination information
type Paginator struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalPages int `json:"total"`
	Data       any `json:"data"`
}

// NewPaginator creates a new Paginator instance
func New(page, pageSize, totalItems int, data any) *Paginator {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	return &Paginator{
		Page:       page,
		PerPage:    pageSize,
		TotalPages: totalPages,
		Data:       data,
	}
}
