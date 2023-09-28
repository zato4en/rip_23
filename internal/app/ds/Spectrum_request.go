package ds

import "gorm.io/gorm"

type Spectrum_request struct {
	gorm.Model
	SatelliteID uint      `json:"-"`
	SpectrumID  uint      `json:"-"`
	Satellite   Satellite `gorm:"foreignKey:ARID" json:"-"`
	Spectrum    Spectrum  `gorm:"foreignKey:SpectrumID" json:"-"`
}
