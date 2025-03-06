package core

import (
	"fmt"
)

type PageRequest struct {
	Sort        string `form:"sort"`
	CurrentPage int    `form:"current_page"`
	PerPage     int    `form:"per_page"`
}

const (
	defaultPerPage = 20
	defaultPage    = 1
	defaultSort    = "id DESC"
)

func (p *PageRequest) GetOffset() int {
	page := defaultPage
	if p.CurrentPage != 0 && p.CurrentPage > 0 {
		page = p.CurrentPage
	}
	return p.GetPerPage() * (page - 1)
}

func (p *PageRequest) GetCurrentPage() int {
	if p.CurrentPage <= 0 {
		return 1
	}
	return p.CurrentPage
}

func (p *PageRequest) GetPerPage() int {
	if p.PerPage <= 0 {
		return defaultPerPage
	}
	return p.PerPage
}

func (p *PageRequest) BuildLimitQuery() string {
	return fmt.Sprintf(" ORDER BY %s LIMIT %d OFFSET %d", p.GetOrderBy(), p.GetPerPage(), p.GetOffset())
}

func (p *PageRequest) GetOrderBy() string {
	if p.Sort == "" {
		return defaultSort
	}
	var field, order string
	if len(p.Sort) > 0 && p.Sort[0] == '-' {
		order = "desc"
		field = p.Sort[1:]
	} else {
		order = "asc"
		field = p.Sort
	}
	if field == "kyb" {
		field = "kyc"
	}

	return fmt.Sprintf("%s %s", field, order)
}

type PageResponse struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	PageCount   int `json:"page_count"`
	TotalCount  int `json:"total_count"`
}

func (p *PageResponse) Set(page, limit, totalData int) {
	p.CurrentPage = page
	p.PerPage = limit
	p.PageCount = (totalData / limit)
	p.TotalCount = totalData

	if (totalData % limit) > 0 {
		p.PageCount++
	}
}
