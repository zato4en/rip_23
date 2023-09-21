package ds

import "gorm.io/gorm"

type Spectrum struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255)" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Min_len     float64 `gorm:"not null" json:"min_len"`
	Max_len     float64 `gorm:"not null" json:"max_len"`
	Image       string  `gorm:"type:varchar(255)" json:"image"`
	IsDelete    bool    `json:"is_delete"`
}
