package helper

import (
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalData  int64       `json:"total_data"`
	TotalPages int         `json:"total_pages"`
	Items      interface{} `json:"items"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB,limit int) func(db *gorm.DB) *gorm.DB {
	if limit == 0 {
		limit = 10
	}
	var totalData int64
	db.Model(value).Count(&totalData)
	pagination.TotalData = totalData
	pagination.Limit = limit
	totalPages := int(math.Ceil(float64(totalData) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.Limit).Order(pagination.GetSort())
	}
}
