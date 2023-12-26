package repository

import (
	"rip2023/internal/app/ds"
	"time"
)

func (r *Repository) SpectrumRequestsList() (*[]ds.Spectrum_request, error) {
	var Spectrum_requests []ds.Spectrum_request
	result := r.db.Select("satellite_id", "Spectrum_id", "Satellite_number").Find(&Spectrum_requests)
	return &Spectrum_requests, result.Error
}

func (r *Repository) UpdateSpectrumNumberInRequest(updatedSpectrumRequest *ds.Spectrum_request) error {
	var oldSpectrumRequest ds.Spectrum_request
	if result := r.db.First(&oldSpectrumRequest, updatedSpectrumRequest.SatelliteID, updatedSpectrumRequest.SpectrumID); result.Error != nil {
		return result.Error
	}
	if updatedSpectrumRequest.Satellite_number != 0 {
		oldSpectrumRequest.Satellite_number = updatedSpectrumRequest.Satellite_number
	}

	*updatedSpectrumRequest = oldSpectrumRequest
	result := r.db.Save(updatedSpectrumRequest)
	return result.Error
}

func (r *Repository) AddSpectrumToRequest(pr *struct {
	SpectrumId uint `json:"Spectrum_id"`
}, userid uint) error {
	var SatelliteRequest ds.Satellite
	var user ds.Users
	r.db.Where("user_id = ? AND status = ?", userid, "черновик").First(&SatelliteRequest)
	r.db.Where("id = ?", userid).First(&user)
	if SatelliteRequest.ID == 0 {
		newRequest := ds.Satellite{UserID: user.ID, UserLogin: user.Login, Status: "черновик", DateCreate: time.Now(), Percentage: "0%"}
		r.db.Create(&newRequest)
		SatelliteRequest = newRequest
	}
	query := "INSERT INTO Spectrum_requests (Satellite_id, Spectrum_id) VALUES ($1,$2) ON CONFLICT DO NOTHING;"
	err := r.db.Exec(query, SatelliteRequest.ID, pr.SpectrumId)
	if err != nil {
		return err.Error
	}
	return nil
}

func (r *Repository) DeleteSpectrumRequest(SatelliteID, SpectrumID uint) error {
	var SpectrumRequest ds.Spectrum_request
	err := r.db.Where("satellite_id = ? AND spectrum_id = ?", SatelliteID, SpectrumID).Delete(&SpectrumRequest).Error
	if err != nil {
		return err
	}
	return nil
}
