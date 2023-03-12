package model

import (
	"time"

	"gorm.io/gorm"
)

type Space struct {
	gorm.Model
	Nama      string   `gorm:"type:VARCHAR(50); NOT NULL" json:"nama"`
	Kategori  string   `gorm:"type:VARCHAR(20); NOT NULL" json:"kategori"`
	Deskripsi string   `gorm:"type:LONGTEXT;" json:"deskripsi"`
	Alamat    string   `gorm:"type:LONGTEXT; NOT NULL" json:"alamat"`
	Harga     int      `json:"harga"`
	Periode   int      `json:"periode"`
	Foto      string   `gorm:"type:LONGTEXT" json:"foto"`
	Rating    float64  `json:"rating"`
	OwnerID   uint     `json:"owner_id"`
	Options   []Option `json:"options"`
}

var Category []string = []string{"private", "coworking", "meeting room"}

type Option struct {
	gorm.Model
	Rentang string `gorm:"type:VARCHAR(10); NOT NULL" json:"rentang"`
	Dates   []Date `json:"dates"`
	SpaceID uint   `json:"space_id"`
}

type Date struct {
	gorm.Model
	Tanggal  time.Time `json:"tanggal"`
	Tersedia bool      `json:"tersedia"`
	SpaceID  uint      `json:"space_id"`
	OptionID uint      `json:"option_id"`
}

type CreateSpaceRequest struct {
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

type CreateDateRequest struct {
}

type CategoryRequest struct {
	Kategori string `json:"nama" form:"kategori"`
	Search   string `json:"search" form:"search"`
}
