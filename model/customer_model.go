package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Email    string  `gorm:"type:VARCHAR(50);NOT NULL;UNIQUE" json:"email"`
	Password string  `gorm:"type:TEXT;NOT NULL" json:"password"`
	Role     string  `gorm:"type:VARCHAR(20); default:'customer'" json:"role"`
	Nama     string  `gorm:"type:VARCHAR(50);NOT NULL" json:"nama"`
	Orders   []Order `json:"orders"`
}

type CreateCustomerRequest struct {
	Nama      string `json:"nama" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Konfirmpw string `json:"konfirmpw" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
