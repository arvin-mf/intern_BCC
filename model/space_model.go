package model

import (
	"gorm.io/gorm"
)

type CreateSpaceRequest struct {
	gorm.Model
	Nama       string `json:"nama" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	Alamat     string `json:"alamat" binding:"required"`
	Harga      int    `json:"harga" binding:"required"`
	Periode    int    `json:"periode" binding:"required"`
	Foto       string `json:"foto"`
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
