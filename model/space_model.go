package model

import (
	"gorm.io/gorm"
)

type CreateSpaceRequest struct {
	gorm.Model
	Nama       string `json:"nama"`
	CategoryID uint   `json:"category_id"`
	Alamat     string `json:"alamat"`
	Harga      int    `json:"harga"`
	Periode    int    `json:"periode"`
	OwnerID    uint   `json:"owner_id"`
}

type CreateOptionRequest struct {
	SpaceID uint   `json:"space_id"`
	Rentang string `json:"rentang"`
}

type CategoryRequest struct {
	Kategori string `json:"nama" form:"kategori"`
}
