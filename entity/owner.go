package entity

import (
	"gorm.io/gorm"
)

type Place struct {
	gorm.Model
	Nama    string   `gorm:"type:VARCHAR(50); NOT NULL" json:"nama"`
	Alamat  string   `gorm:"type:LONGTEXT; NOT NULL" json:"alamat"`
	OwnerID uint     `json:"owner_id"`
	Reviews []Review `json:"reviews"`
}

type Owner struct {
	gorm.Model
	Email    string  `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password string  `gorm:"type:TEXT; NOT NULL" json:"password"`
	Nama     string  `gorm:"type:VARCHAR(50)" json:"nama"`
	Whatsapp string  `gorm:"type:VARCHAR(15); NOT NULL" json:"wnatsapp"`
	Spaces   []Space `json:"spaces"`
}
