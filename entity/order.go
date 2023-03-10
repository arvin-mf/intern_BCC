package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint   `json:"customer_id"`
	SpaceID    uint   `json:"space_id"`
	Review     Review `json:"review"`
}
