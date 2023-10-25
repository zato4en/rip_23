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
	Satellite         string             `gorm:"type:varchar(255)" json:"satellite"`
	UserID            uint               `json:"user_id"`
	ModerID           uint               `json:"-"`
	Spectrum_requests []Spectrum_request `json:"spectrum_requests" gorm:"foreignkey:SatelliteID"`
	User              Users              `gorm:"foreignKey:UserID" json:"-"`
}
