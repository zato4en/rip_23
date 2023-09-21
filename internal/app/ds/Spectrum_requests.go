package ds

import "gorm.io/gorm"

type SpectrumRequest struct {
	gorm.Model
	ARID            uint            `json:"-"`
	SpectrumID      uint            `json:"-"`
	AnalysisRequest AnalysisRequest `gorm:"foreignKey:ARID" json:"-"`
	Spectrum        Spectrum        `gorm:"foreignKey:SpectrumID" json:"-"`
}
