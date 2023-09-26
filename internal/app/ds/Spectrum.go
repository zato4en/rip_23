package ds

import "gorm.io/gorm"

type Spectrum struct {
	gorm.Model
	Freq        float64 `gorm:"not null" json:"freq"`
	Len         float64 `gorm:"not null" json:"len"`
	Description string  `gorm:"type:text" json:"description"`
	Image       string  `gorm:"type:varchar(255)" json:"image"`
	IsDelete    bool    `json:"is_delete"`
}
