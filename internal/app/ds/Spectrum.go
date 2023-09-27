package ds

import "gorm.io/gorm"

type Spectrum struct {
	gorm.Model
	Len  float64 `gorm:"not null" json:"len"`
	Freq float64 `gorm:"not null" json:"freq"`

	Description string `gorm:"type:text" json:"description"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
	IsDelete    bool   `json:"is_delete"`
}
