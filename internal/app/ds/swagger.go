package ds

type SatellitesListRes struct {
	Status    string      `json:"status"`
	Satellite []Satellite `json:"Satellite"`
}

type SatellitesListRes2 struct {
	Status    string      `json:"status"`
	Satellite []Satellite `json:"Satellite"`
}

type DeleteSpectrumInRequestReq struct {
	ID int `json:"id"`
}

type DeleteSpectrumInRequestRes struct {
	Status                   string `json:"status"`
	DeletedSpectrumInRequest int    `json:"deleted_Spectrum_in_request"`
}

type UpdateSpectrumInRequestNumberReq struct {
	SpectrumInRequestID int `json:"id"`
	SatelliteNumber     int `json:"Satellite_number"`
}

type UpdateSpectrumInRequestNumberRes struct {
	Status string `json:"status"`
	ID     uint   `json:"id"`
}

type DeleteSpectrumRes struct {
	DeletedId int `json:"deleted_id"`
}

//type AddImageRes struct {
//	Status   string `json:"status"`
//	ImageUrl string `json:"image_url"`
//}

type UpdatedSatelliteRes struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	DateCreate     string `json:"date_create"`
	DateFormation  string `json:"date_formation"`
	DateCompletion string `json:"date_completion"`
	Status         string `json:"status"`
	AMS            string `json:"ams"`
	UserID         uint   `json:"user_id"`
	ModerID        uint   `json:"moder_id"`
	UserLogin      string `json:"user_login"`
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
	SatelliteID uint `json:"id"`
	//Status   string `json:"status"`
}

type UpdateStatusForUserReq struct {
	SatelliteID uint `json:"id"`
	//Status   string `json:"status" example:"в работе"`
}

type DeleteSpectrumReq struct {
	ID string `json:"id"`
}

type UpdateSpectrumReq struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Status      string `json:"status"`
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
	ID   string `json:"id"`
	Name string `json:"name"`
	//Status      string `json:"status"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type AddSpectrumResp struct {
	//Status   string `json:"status"`
	SpectrumId string `json:"Spectrum_id"`
}

type SpectrumsListResp struct {
	//Status  string   `json:"status"`
	Spectrums []Spectrum `json:"Spectrums"`
	//BasketId string `json:"basket_id"`
}
