package entity

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
	OwnerID   uint     `json:"owner_id"`
	Options   []Option `json:"options"`
}

var Category []string = []string{"private", "coworking", "meeting room"}

type Option struct {
	gorm.Model
	Rentang string `gorm:"type:VARCHAR(10); NOT NULL" json:"rentang"`
	Dates   []Date `json:"dates"`
	OwnerID uint   `json:"owner_id"`
	SpaceID uint   `json:"space_id"`
}

type Date struct {
	gorm.Model
	Tanggal  time.Time `json:"tanggal"`
	Tersedia bool      `json:"tersedia"`
	OwnerID  uint      `json:"owner_id"`
	SpaceID  uint      `json:"space_id"`
	OptionID uint      `json:"option_id"`
}
