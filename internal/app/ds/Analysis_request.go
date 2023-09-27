package ds

import (
	"gorm.io/gorm"
	"time"
)

type AnalysisRequest struct {
	gorm.Model
	DateSend  time.Time `json:"date_send"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	Satellite string    `gorm:"type:varchar(255)" json:"satellite"`
	UserID    uint      `json:"-"`
	ModerID   uint      `json:"-"`
	User      Users     `gorm:"foreignKey:UserID" json:"-"`
}
