package helper

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type PaginationMeta struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

func PaginateScope(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(pagination.Limit).Offset(GetOffset(pagination.Page, pagination.Limit))
	}
}

func GetOffset(page, limit int) int {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	return (page - 1) * limit
}

func GetTotalPage(total int64, limit int) int {
	return (int(total) + limit - 1) / limit
}
