package repository

import (
	"rip2023/internal/app/ds"
	"rip2023/internal/app/utils"
)

func (r *Repository) SatellitesList(userID, datestart, dateend, status string) (*[]ds.Satellite, error) {
	var Satellites []ds.Satellite
	db := r.db.Preload("User").Where("status !=?", utils.DeletedString)

	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}

	if datestart != "" && dateend != "" {
		db = db.Where("date_formation > ? AND date_formation < ?", datestart, dateend)
	}

	if status != "" {
		db = db.Where("status = ?", status)
	}
	for i := range Satellites {
		Satellites[i].UserLogin = Satellites[i].User.Login
	}
	result := db.Find(&Satellites)
	return &Satellites, result.Error
}

func (r *Repository) SatelliteById(id string) (*ds.Satellite, error) {
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
func (r *Repository) UsersSatellite() (*[]ds.Satellite, error) {
	var satellite []ds.Satellite
	result := r.db.Preload("User").Preload("Spectrum_requests.Spectrum").Where("user_id = ?", 1).Find(&satellite)
	return &satellite, result.Error
}
func (r *Repository) SatellitesListByDate(datestart, dateend string) (*[]ds.Satellite, error) {
	var Satellites []ds.Satellite
	result := r.db.Preload("User").Where("date_formed > ? AND date_formed < ?", datestart, dateend).Find(&Satellites)
	return &Satellites, result.Error
}

func (r *Repository) SatellitesListByUser(id uint) (*[]ds.Satellite, error) {
	var satellites []ds.Satellite
	result := r.db.Preload("User").Where("user_id = ?", id).Find(&satellites)
	return &satellites, result.Error
}
func (r *Repository) SatellitesListByStatus(status string) (*[]ds.Satellite, error) {
	var satellites []ds.Satellite
	result := r.db.Preload("User").Where("status = ?", status).Find(&satellites)
	return &satellites, result.Error
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
		if updatedSatellite.Status == "в работе" && oldSatellite.Status == "создан" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "завёршён" && oldSatellite.Status == "в работе" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "удалён" && oldSatellite.Status == "отменён" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "отменён" && oldSatellite.Status != "удалён" {
			oldSatellite.Status = updatedSatellite.Status
		}
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
		if updatedSatellite.Status == "в работе" && oldSatellite.Status == "создан" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "завёршён" && oldSatellite.Status == "в работе" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "удалён" && oldSatellite.Status == "отменён" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "отменён" && oldSatellite.Status != "удалён" {
			oldSatellite.Status = updatedSatellite.Status
		}

	}
	*updatedSatellite = oldSatellite
	result := r.db.Save(updatedSatellite)
	return result.Error
}

func (r *Repository) UserUpdateSatelliteStatusById(id int) (*ds.Satellite, error) {
	var Satellite ds.Satellite
	result := r.db.First(&Satellite, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Меняем статус тут
	if Satellite.Status == "создан" {
		Satellite.Status = "в работе"
	} else if Satellite.Status == "в работе" {
		Satellite.Status = "отменён"
	}

	// Сохраняем изменения в базе данных
	if err := r.db.Save(&Satellite).Error; err != nil {
		return nil, err
	}

	return &Satellite, nil
}
func (r *Repository) ModerUpdateSatelliteStatusById(id int) (*ds.Satellite, error) {
	var Satellite ds.Satellite
	result := r.db.First(&Satellite, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Меняем статус тут
	if Satellite.Status == "отменён" {
		Satellite.Status = "удалён"
	}

	// Сохраняем изменения в базе данных
	if err := r.db.Save(&Satellite).Error; err != nil {
		return nil, err
	}

	return &Satellite, nil
}
