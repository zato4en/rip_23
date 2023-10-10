package ds

type Spectrum_request struct {
	SatelliteID      uint `json:"satellite_id" gorm:"primaryKey;auto_increment:false"`
	SpectrumID       uint `json:"spectrum_id" gorm:"primaryKey;auto_increment:false"`
	Satellite_number uint `json:"satellite_number"`
}
