package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Email    string  `gorm:"type:VARCHAR(50);NOT NULL" json:"email"`
	Password string  `gorm:"type:VARCHAR(20);NOT NULL" json:"password"`
	Nama     string  `gorm:"type:VARCHAR(50);NOT NULL" json:"nama"`
	Orders   []Order `json:"orders"`
}
