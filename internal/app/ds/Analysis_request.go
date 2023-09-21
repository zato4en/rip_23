package ds

import (
	"gorm.io/gorm"
	"time"
)

type AnalysisRequest struct {
	gorm.Model
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	Satellite string    `gorm:"type:varchar(255)" json:"satellite"`
	UserID    uint      `json:"-"`
	User      Users     `gorm:"foreignKey:UserID" json:"-"`
}
