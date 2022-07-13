package models

import "gorm.io/gorm"

type WorkingEcosystem struct {
	gorm.Model

	WeName        string `gorm:"not null" json:"we_name"`
	WeDescription string `gorm:"not null" json:"we_description"`
	GlobersId     uint   `json:"glober_id"`
}
