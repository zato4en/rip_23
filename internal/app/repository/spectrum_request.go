package repository

import (
	"rip2023/internal/app/ds"
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

// Тут запрос на SQL потому что если через ГОРМ то будет возвращать айди, а у нас нет поля айди в М-М, поэтому через SQL
// как сделать через ГОРМ тут я без понятия
// Если у нас не находит заявку с айди которое есть в М-М, то мы создаем заявку с таким айди
func (r *Repository) AddSpectrumToRequest(pr *ds.Spectrum_request) error {

	var SatelliteRequest ds.Satellite

	if err := r.db.First(&SatelliteRequest, pr.SatelliteID).Error; err != nil {

		SatelliteRequest = ds.Satellite{ID: pr.SatelliteID, UserID: 1, Status: "существует"}
		r.db.Create(&SatelliteRequest)
	}

	query := "INSERT INTO Spectrum_requests (Satellite_id, Spectrum_id, Satellite_number) VALUES ($1,$2,$3);"

	err := r.db.Exec(query, pr.SatelliteID, pr.SpectrumID, pr.Satellite_number)
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
