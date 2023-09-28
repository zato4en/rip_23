package ds

import "gorm.io/gorm"

type Spectrum_request struct {
	gorm.Model
	ARID       uint      `json:"-"`
	SpectrumID uint      `json:"-"`
	Satellite  Satellite `gorm:"foreignKey:ARID" json:"-"`
	Spectrum   Spectrum  `gorm:"foreignKey:SpectrumID" json:"-"`
}
