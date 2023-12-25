package ds

import (
	"time"
)

type Satellite struct {
	ID                uint               `json:"id" gorm:"primaryKey"`
	DateCreate        time.Time          `json:"date_create"`
	DateFormation     time.Time          `json:"date_formation"`
	DateCompletion    time.Time          `json:"date_completion"`
	Status            string             `gorm:"type:varchar(255)" json:"status"`
	Percentage        string             `gorm:"type:varchar(255)" json:"percentage"`
	Satellite         string             `gorm:"type:varchar(255)" json:"satellite"`
	UserID            uint               `json:"user_id"`
	ModerID           uint               `json:"moder_id"`
	ModerLogin        string             `json:"moder_login"`
	UserLogin         string             `json:"user_login"`
	Spectrum_requests []Spectrum_request `json:"spectrum_requests" gorm:"foreignkey:SatelliteID"`
	User              Users              `gorm:"foreignKey:UserID" json:"-"`
}
