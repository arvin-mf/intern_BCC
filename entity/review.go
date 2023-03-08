package entity

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	CustomerID uint   `json:"customer_id"`
	SpaceID    uint   `json:"space_id"`
	OrderID    uint   `json:"order_id"`
	Ulasan     string `gorm:"type:LONGTEXT" json:"ulasan"`
}
