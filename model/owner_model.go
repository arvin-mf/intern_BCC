package model

import (
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Email    string    `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password string    `gorm:"type:TEXT; NOT NULL" json:"password"`
	Role     string    `gorm:"type:VARCHAR(20); default:'owner'" json:"role"`
	Nama     string    `gorm:"type:VARCHAR(50)" json:"nama"`
	Whatsapp string    `gorm:"type:VARCHAR(15); NOT NULL" json:"wnatsapp"`
	Pictures []Picture `json:"pictures"`
	Spaces   []Space   `json:"spaces"`
}

type Picture struct {
	gorm.Model
	Link    string `gorm:"type:TEXT; NOT NULL" json:"link"`
	OwnerID uint   `json:"owner_id"`
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

type GalleryRequest struct {
	OwnerID uint   `json:"owner_id" binding:"required"`
	Link    string `json:"link" binding:"required"`
}

type DescriptionRequest struct {
	Deskripsi string `json:"deskripsi"`
}

type CapacityRequest struct {
	Kapasitas int `json:"kapasitas"`
}
