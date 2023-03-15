package model

import (
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Email    string  `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password string  `gorm:"type:TEXT; NOT NULL" json:"password"`
	Role     string  `gorm:"type:VARCHAR(20); default:'owner'" json:"role"`
	Nama     string  `gorm:"type:VARCHAR(50)" json:"nama"`
	Whatsapp string  `gorm:"type:VARCHAR(15); NOT NULL" json:"wnatsapp"`
	Spaces   []Space `json:"spaces"`
}

type CreateOwnerRequest struct {
	Email    string `json:"email" binding:"required"`
	Whatsapp string `json:"whatsapp" binding:"required"`
}

type GetByCatRequest struct {
	Kategori int `uri:"kategori" binding:"required"`
}

type AddFacilRequest struct {
	Fasil []string `json:"fasil"`
}

type DescriptionRequest struct {
	Deskripsi string `json:"deskripsi"`
}

type CapacityRequest struct {
	Kapasitas int `json:"kapasitas"`
}
