package domain

import "math"

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

// GetOffset calculates how many records GORM needs to skip
func (p *Pagination) GetOffset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10 // default limit
	}
	return int((p.Page - 1) * p.Limit)
}

// GetLimit returns the limit as an int for GORM
func (p *Pagination) GetLimit() int {
	if p.Limit < 1 {
		p.Limit = 10
	}
	return int(p.Limit)
}

// CalculatePages sets the total pages based on total items found
func (p *Pagination) CalculatePages(totalItems int64) {
	p.TotalItems = totalItems
	p.TotalPages = int64(math.Ceil(float64(totalItems) / float64(p.GetLimit())))
}