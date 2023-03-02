package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Waktu      int    `json:"waktu"`
	CustomerID uint   `json:"customer_id"`
	SpaceID    uint   `json:"space_id"`
	Review     Review `json:"review"`
}
