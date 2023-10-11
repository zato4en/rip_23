package repository

import (
	"rip2023/internal/app/ds"
	"rip2023/internal/app/utils"
)

func (r *Repository) SatellitesList() (*[]ds.Satellite, error) {
	var Satellites []ds.Satellite
	result := r.db.Preload("User").Where("status =?", utils.ExistsString).Find(&Satellites)
	return &Satellites, result.Error
}

func (r *Repository) SatelliteById(id uint) (*ds.Satellite, error) {
	Satellite := ds.Satellite{}
	result := r.db.Preload("User").Preload("Spectrum_requests.Spectrum").First(&Satellite, id)
	return &Satellite, result.Error
}

func (r *Repository) DeleteSatellite(id uint) error {

	err := r.db.Model(&ds.Satellite{}).Where("id = ?", id).Update("status", utils.DeletedString)
	if err != nil {
		return err.Error
	}
	return nil

}

func (r *Repository) UpdateSatellite(updatedSatellite *ds.Satellite) error {
	oldSatellite := ds.Satellite{}
	if result := r.db.First(&oldSatellite, updatedSatellite.ID); result.Error != nil {
		return result.Error
	}
	if updatedSatellite.DateCreated.String() != utils.EmptyDate {
		oldSatellite.DateCreated = updatedSatellite.DateCreated
	}
	if updatedSatellite.DateFormed.String() != utils.EmptyDate {
		oldSatellite.DateFormed = updatedSatellite.DateFormed
	}
	if updatedSatellite.DateAccepted.String() != utils.EmptyDate {
		oldSatellite.DateAccepted = updatedSatellite.DateAccepted
	}
	if updatedSatellite.Status != "" {
		oldSatellite.Status = updatedSatellite.Status
	}
	if updatedSatellite.Satellite != "" {
		oldSatellite.Satellite = updatedSatellite.Satellite
	}
	if updatedSatellite.UserID != utils.EmptyInt {
		oldSatellite.UserID = updatedSatellite.UserID
	}
	if updatedSatellite.ModerID != utils.EmptyInt {
		oldSatellite.ModerID = updatedSatellite.ModerID
	}
	*updatedSatellite = oldSatellite
	result := r.db.Save(updatedSatellite)
	return result.Error
}

func (r *Repository) UpdateSatelliteStatus(updatedSatellite *ds.Satellite) error {
	oldSatellite := ds.Satellite{}
	if result := r.db.First(&oldSatellite, updatedSatellite.ID); result.Error != nil {
		return result.Error
	}
	if updatedSatellite.Status != "" {
		oldSatellite.Status = updatedSatellite.Status
	}
	*updatedSatellite = oldSatellite
	result := r.db.Save(updatedSatellite)
	return result.Error
}
