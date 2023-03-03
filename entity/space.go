package entity

import (
	"gorm.io/gorm"
)

type Space struct {
	gorm.Model
	Nama     string `gorm:"type:VARCHAR(50); NOT NULL" json:"nama"`
	Kategori string `gorm:"type:VARCHAR(50); NOT NULL" json:"kategori"`
	// Opsi     []int  `json:"opsi"`
	Per     int  `json:"per"`
	Harga   int  `json:"harga"`
	PlaceID uint `json:"place_id"`
}
