package pagination

import (
	"math"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int   `json:"limit,omitempty"`
	Page       int   `json:"page,omitempty"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

func NewPagination(page, limit string) *Pagination {
	p := &Pagination{}
	
	if page == "" {
		p.Page = 1
	} else {
		pageNum, _ := strconv.Atoi(page)
		if pageNum <= 0 {
			pageNum = 1
		}
		p.Page = pageNum
	}

	if limit == "" {
		p.Limit = 10
	} else {
		limitNum, _ := strconv.Atoi(limit)
		switch {
		case limitNum > 100:
			p.Limit = 100
		case limitNum <= 0:
			p.Limit = 10
		default:
			p.Limit = limitNum
		}
	}

	return p
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) GetLimit() int {
	return p.Limit
}

func (p *Pagination) SetTotalRows(totalRows int64) {
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.Limit)))
}

// Paginate adds pagination to GORM query
func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit())
	}
}