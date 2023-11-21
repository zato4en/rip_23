package ds

type SatellitesListRes struct {
	Status    string      `json:"status"`
	Satellite []Satellite `json:"Satellite"`
}

type SatellitesListRes2 struct {
	Status    string      `json:"status"`
	Satellite []Satellite `json:"Satellite"`
}

type DeleteSpectrumRes struct {
	DeletedId int `json:"deleted_id"`
}

//type AddImageRes struct {
//	Status   string `json:"status"`
//	ImageUrl string `json:"image_url"`
//}

type UpdatedSatelliteRes struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	DateCreated  string `json:"date_created"`
	DateFormed   string `json:"date_formed"`
	DateAccepted string `json:"date_accepted"`
	Status       string `gorm:"type:varchar(255)" json:"status"`
	Satellite    string `gorm:"type:varchar(255)" json:"satellite"`
	UserID       uint   `json:"user_id"`
	ModerID      uint   `json:"moder_id"`
	UserLogin    string `json:"user_login"`
}

type DeleteSatelliteRes struct {
	Status      string `json:"status"`
	SatelliteId uint   `json:"Satellite_id"`
}

type DeleteSatelliteReq struct {
	ID uint `json:"id"`
}

type UpdateSatelliteReq struct {
	ID          uint   `json:"id"`
	AMS         string `json:"ams"`
	Description string `json:"description"`
}

type UpdateStatusForModeratorReq struct {
	SatelliteID uint   `json:"Satellite_id"`
	Status      string `json:"status"`
}

type UpdateStatusForUserReq struct {
	Status string `json:"status" example:"2"`
}

type DeleteSpectrumReq struct {
	ID string `json:"id"`
}

type UpdateSpectrumReq struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"city_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type AddSpectrumToRequestReq struct {
	SpectrumID int `json:"Spectrum_id" binding:"required" example:"1"`
	//SerialNumber int `json:"serial_number" binding:"required" example:"1"`
}

type AddSpectrumToRequestResp struct {
	Status string `json:"status"`
	Id     int    `json:"id"`
}

type UpdateSpectrumResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type AddSpectrumResp struct {
	Status     string `json:"status"`
	SpectrumId string `json:"Spectrum_id"`
}

type SpectrumsListResp struct {
	Status    string     `json:"status"`
	Spectrums []Spectrum `json:"Spectrums"`
	//BasketId string `json:"basket_id"`
}

//type AddPlaIntoHikeRequest struct {
//}
