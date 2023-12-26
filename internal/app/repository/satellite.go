package repository

import (
	"rip2023/internal/app/ds"
	"rip2023/internal/app/utils"
	"time"
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
func (r *Repository) UsersSatellite(userid uint) (*[]ds.Satellite, error) {
	var satellite []ds.Satellite
	result := r.db.Preload("User").Preload("Spectrum_requests.Spectrum").Where("user_id = ? and status != ?", userid, "черновик").Find(&satellite)
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
	if updatedSatellite.DateCreate.String() != utils.EmptyDate {
		oldSatellite.DateCreate = updatedSatellite.DateCreate
	}
	if updatedSatellite.DateFormation.String() != utils.EmptyDate {
		oldSatellite.DateFormation = updatedSatellite.DateFormation
	}
	if updatedSatellite.DateCompletion.String() != utils.EmptyDate {
		oldSatellite.DateCompletion = updatedSatellite.DateCompletion
	}

	if updatedSatellite.Status != "" {
		if updatedSatellite.Status == "в работе" && oldSatellite.Status == "черновик" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "завершен" && oldSatellite.Status == "в работе" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "удален" && oldSatellite.Status == "отменен" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "отменен" && oldSatellite.Status != "удален" {
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

func (r *Repository) UpdateSatelliteAsyncStatus(satelliteID int, percentage string) error {
	// Поиск существующего объекта Satellite по ID
	existingSatellite := ds.Satellite{}
	iduint := uint(satelliteID)
	if result := r.db.First(&existingSatellite, iduint); result.Error != nil {
		return result.Error
	}

	// Обновление поля asyncStatus в найденной записи
	existingSatellite.Percentage = percentage

	// Сохранение изменений в базу данных
	result := r.db.Save(&existingSatellite)
	return result.Error
}

func (r *Repository) UpdateSatelliteStatus(updatedSatellite *ds.Satellite) error {
	oldSatellite := ds.Satellite{}
	if result := r.db.First(&oldSatellite, updatedSatellite.ID); result.Error != nil {
		return result.Error
	}
	if updatedSatellite.Status != "" {
		if updatedSatellite.Status == "в работе" && oldSatellite.Status == "черновик" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "завершен" && oldSatellite.Status == "в работе" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "удален" && oldSatellite.Status == "отменен" {
			oldSatellite.Status = updatedSatellite.Status
		}
		if updatedSatellite.Status == "отменен" && oldSatellite.Status != "удален" {
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
	if Satellite.Status == "черновик" {
		Satellite.Status = "в работе"
		Satellite.DateFormation = time.Now()
	} else if Satellite.Status == "в работе" {
		Satellite.Status = "отменен"
	}

	// Сохраняем изменения в базе данных
	if err := r.db.Save(&Satellite).Error; err != nil {
		return nil, err
	}

	return &Satellite, nil
}
func (r *Repository) ModerUpdateSatelliteStatusById(id int, modername string, status string) (*ds.Satellite, error) {
	var Satellite ds.Satellite
	var user ds.Users
	r.db.Where("user_name = ?", modername).First(&user)

	result := r.db.First(&Satellite, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Меняем статус тут
	Satellite.Status = status
	Satellite.ModerID = user.ID
	Satellite.ModerLogin = modername

	// Сохраняем изменения в базе данных
	if err := r.db.Save(&Satellite).Error; err != nil {
		return nil, err
	}

	return &Satellite, nil
}

func (r *Repository) UsersUpdateSatellite(updatedSatellite *ds.Satellite, userid uint) error {
	oldSatellite := ds.Satellite{}
	result := r.db.Where("user_id = ?", userid).Find(&oldSatellite)
	if result.Error != nil {
		return result.Error
	}
	if updatedSatellite.DateCreate.String() != utils.EmptyDate {
		oldSatellite.DateCreate = updatedSatellite.DateCreate
	}
	if updatedSatellite.DateFormation.String() != utils.EmptyDate {
		oldSatellite.DateFormation = updatedSatellite.DateFormation
	}
	if updatedSatellite.DateCompletion.String() != utils.EmptyDate {
		oldSatellite.DateCompletion = updatedSatellite.DateCompletion
	}
	if updatedSatellite.Status != "" {
		if updatedSatellite.Status == "в работе" && oldSatellite.Status == "черновик" {
			oldSatellite.Status = updatedSatellite.Status
		} else if updatedSatellite.Status == "отменен" && oldSatellite.Status == "в работе" {
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
	result = r.db.Save(updatedSatellite)
	return result.Error
}
