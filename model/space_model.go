package model

import (
	"time"

	"gorm.io/gorm"
)

type Space struct {
	gorm.Model
	Nama       string     `gorm:"type:VARCHAR(50); NOT NULL" json:"nama"`
	Kategori   string     `gorm:"type:VARCHAR(20); NOT NULL" json:"kategori"`
	Deskripsi  string     `gorm:"type:LONGTEXT" json:"deskripsi"`
	Alamat     string     `gorm:"type:LONGTEXT NOT NULL" json:"alamat"`
	Facilities []Facility `json:"facilities"`
	Kapasitas  int        `json:"kapasitas"`
	Harga      int        `json:"harga"`
	Periode    int        `json:"periode"`
	Foto       string     `gorm:"type:LONGTEXT" json:"foto"`
	Lat        float64    `json:"lat"`
	Lon        float64    `json:"lon"`
	Jarak      float64    `json:"jarak"`
	Rating     float64    `gorm:"default:5.0" json:"rating"`
	OwnerID    uint       `json:"owner_id"`
	Options    []Option   `json:"options"`
}

var Category []string = []string{"private", "coworking", "meeting room"}

type Facility struct {
	gorm.Model
	Ket     string `gorm:"type:TEXT" json:"ket"`
	SpaceID uint   `json:"space_id"`
}

type Option struct {
	gorm.Model
	Rentang string `gorm:"type:VARCHAR(10); NOT NULL" json:"rentang"`
	Dates   []Date `json:"dates"`
	SpaceID uint   `json:"space_id"`
}

type Date struct {
	gorm.Model
	Tanggal  time.Time `gorm:"type:DATE" json:"tanggal"`
	Tersedia bool      `gorm:"default:true" json:"tersedia"`
	OptionID uint      `json:"option_id"`
	Hari     string    `gorm:"type:VARCHAR(10); NOT NULL" json:"hari"`
	Tgl      string    `gorm:"type:VARCHAR(20); NOT NULL" json:"tgl"`
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
	OptionID uint   `json:"option_id" binding:"required"`
	Tgl      string `json:"tgl" binding:"required"`
	Hari     string `json:"hari" binding:"required"`
}

type InputLocation struct {
	ID  uint    `json:"id" binding:"required"`
	Lat float64 `json:"lat" binding:"required"`
	Lon float64 `json:"lon" binding:"required"`
}

type CategoryRequest struct {
	Kategori string `json:"nama" form:"kategori"`
	Search   string `json:"search" form:"search"`
}

type PictureRequest struct {
	SpaceID uint   `json:"space_id" binding:"required"`
	Link    string `json:"link" binding:"required"`
}

type UserLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
