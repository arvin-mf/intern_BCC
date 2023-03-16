package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint   `json:"customer_id"`
	SpaceID    uint   `json:"space_id"`
	Review     Review `json:"review"`
}

type Review struct {
	gorm.Model
	CustomerID uint   `json:"customer_id"`
	SpaceID    uint   `json:"space_id"`
	OrderID    uint   `json:"order_id"`
	Nama       string `gorm:"type:VARCHAR(50)" json:"nama"`
	Ulasan     string `gorm:"type:LONGTEXT" json:"ulasan"`
	Rating     int    `json:"rating"`
}

type CreateReviewRequest struct {
	Ulasan string `json:"ulasan" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
}
