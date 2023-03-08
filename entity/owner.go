package entity

import (
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Email    string  `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password string  `gorm:"type:TEXT; NOT NULL" json:"password"`
	Nama     string  `gorm:"type:VARCHAR(50)" json:"nama"`
	Whatsapp string  `gorm:"type:VARCHAR(15); NOT NULL" json:"wnatsapp"`
	Spaces   []Space `json:"spaces"`
}
