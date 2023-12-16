package ds

import (
	"time"
)

type Satellite struct {
	ID                uint               `json:"id" gorm:"primaryKey"`
	DateCreated       time.Time          `json:"date_created"`
	DateFormed        time.Time          `json:"date_formed"`
	DateAccepted      time.Time          `json:"date_accepted"`
	Status            string             `gorm:"type:varchar(255)" json:"status"`
	Percentage        string             `gorm:"type:varchar(255)" json:"percentage"`
	Satellite         string             `gorm:"type:varchar(255)" json:"satellite"`
	UserID            uint               `json:"user_id"`
	ModerID           uint               `json:"moder_id"`
	UserLogin         string             `json:"user_login"`
	Spectrum_requests []Spectrum_request `json:"spectrum_requests" gorm:"foreignkey:SatelliteID"`
	User              Users              `gorm:"foreignKey:UserID" json:"-"`
}
