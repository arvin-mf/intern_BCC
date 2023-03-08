package model

import (
	"math"

	"gorm.io/gorm"
)

type CreateSpaceRequest struct {
	gorm.Model
	Nama       string `json:"nama" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	Alamat     string `json:"alamat" binding:"required"`
	Harga      int    `json:"harga" binding:"required"`
	Periode    int    `json:"periode" binding:"required"`
	OwnerID    uint   `json:"owner_id" binding:"required"`
}

type CreateOptionRequest struct {
	SpaceID uint   `json:"space_id" binding:"required"`
	Rentang string `json:"rentang" binding:"required"`
}

type CategoryRequest struct {
	Kategori string `json:"nama" form:"kategori"`
	Search   string `json:"search" form:"search"`
}

type PaginParam struct {
	Limit           int `json:"limit" form:"limit"`
	Page            int `json:"page" form:"page"`
	Offset          int `json:"offset"`
	TotalElements   int `json:"total_elements"`
	CurrentElements int `json:"current_elements"`
	TotalPages      int `json:"total_pages"`
	CurrentPage     int `json:"current_page"`
}

func (pp *PaginParam) FormatPagin() {
	if pp.Limit == 0 {
		pp.Limit = 9
	}
	if pp.Page == 0 {
		pp.Page = 1
	}
	pp.Offset = (pp.Page - 1) * pp.Limit
}

func (pp *PaginParam) ProcessPagin(totalElements int) {
	pp.TotalElements = totalElements
	pp.TotalPages = int(math.Ceil(float64(pp.TotalElements) / float64(pp.Limit)))
	pp.CurrentPage = pp.Page
	if pp.TotalPages > pp.TotalElements/pp.Limit {
		if pp.Page < pp.TotalPages {
			pp.CurrentElements = pp.Limit
		} else {
			pp.CurrentElements = pp.TotalElements - (pp.TotalPages-1)*pp.Limit
		}
	}
}
