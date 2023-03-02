package entity

import (
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Email    string  `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password string  `gorm:"type:VARCHAR(20); NOT NULL" json:"password"`
	Nama     string  `gorm:"type:VARCHAR(50)" json:"nama"`
	Whatsapp string  `gorm:"type:VARCHAR(15); NOT NULL" json:"wnatsapp"`
	Places   []Place `json:"places"`
}

type Place struct {
	gorm.Model
	Nama    string   `gorm:"type:VARCHAR(50); NOT NULL" json:"nama"`
	Alamat  string   `gorm:"type:LONGTEXT; NOT NULL" json:"alamat"`
	OwnerID uint     `json:"owner_id"`
	Spaces  []Space  `json:"spaces"`
	Reviews []Review `json:"reviews"`
}

type Space struct {
	gorm.Model
	Nama     string `gorm:"type:VARCHAR(50); NOT NULL" json:"nama"`
	Kategori string `gorm:"type:VARCHAR(50); NOT NULL" json:"kategori"`
	// Opsi     []int  `json:"opsi"`
	Per     int  `json:"per"`
	Harga   int  `json:"harga"`
	PlaceID uint `json:"place_id"`
}
