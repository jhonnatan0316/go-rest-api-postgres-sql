package models

import "gorm.io/gorm"

type Globers struct {
	gorm.Model

	GloberName        string             `gorm:"not null" json:"glober_name"`
	GloberMail        string             `gorm:"not null;unique_index" json:"glober_mail"`
	WorkingEcosystems []WorkingEcosystem `json:"working_ecosystem"`
}
