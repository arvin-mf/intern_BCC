package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint      `json:"customer_id"`
	SpaceID    uint      `json:"space_id"`
	Waktu      string    `json:"waktu"`
	Tanggal    time.Time `json:"tanggal"`
	Review     Review    `json:"review"`
}
