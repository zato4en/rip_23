package repository

import (
	"rip2023/internal/app/ds"
	"strings"
)

func (r *Repository) SpectrumsList() (*[]ds.Spectrum, error) {
	var Spectrums []ds.Spectrum
	r.db.Where("is_delete = ?", false).Find(&Spectrums)
	return &Spectrums, nil
}

func (r *Repository) SearchSpectrum(search string) (*[]ds.Spectrum, error) {
	var Spectrums []ds.Spectrum
	r.db.Find(&Spectrums)

	var filteredSpectrums []ds.Spectrum
	for _, Spectrum := range Spectrums {
		if strings.Contains(strings.ToLower(Spectrum.Name), strings.ToLower(search)) {
			filteredSpectrums = append(filteredSpectrums, Spectrum)
		}
	}

	return &filteredSpectrums, nil
}

func (r *Repository) SpectrumById(id int) (*ds.Spectrum, error) {
	var spectrums ds.Spectrum
	r.db.Find(&spectrums, id)
	return &spectrums, nil
}

func (r *Repository) DeleteSpectrum(id uint) error {
	//query := "UPDATE Spectrums SET is_delete = true WHERE id = $1"
	//r.db.Exec(query, id)
	err := r.db.Model(&ds.Spectrum{}).Where("id = ?", id).Update("is_delete", true)
	if err != nil {
		return err.Error
	}
	return nil
}
func (r *Repository) AddSpectrum(Spectrum *ds.Spectrum) error {
	result := r.db.Create(&Spectrum)
	return result.Error
}
func (r *Repository) UpdateSpectrum(updatedSpectrum *ds.Spectrum) error {
	var oldSpectrum ds.Spectrum
	if result := r.db.First(&oldSpectrum, updatedSpectrum.ID); result.Error != nil {
		return result.Error
	}
	if updatedSpectrum.Description != "" {
		oldSpectrum.Description = updatedSpectrum.Description
	}
	if updatedSpectrum.Name != "" {
		oldSpectrum.Name = updatedSpectrum.Name
	}
	if updatedSpectrum.Len != 0 {
		oldSpectrum.Len = updatedSpectrum.Len
	}
	if updatedSpectrum.Freq != 0 {
		oldSpectrum.Freq = updatedSpectrum.Freq
	}

	//if updatedSpectrum.Image != "" {
	//	oldSpectrum.Image = updatedSpectrum.Image
	//}

	oldSpectrum.IsDelete = updatedSpectrum.IsDelete

	*updatedSpectrum = oldSpectrum
	result := r.db.Save(updatedSpectrum)
	return result.Error
}

func (r *Repository) UpdateSpectrumImage(id string, newImageURL string) error {
	spectrum := ds.Spectrum{}
	if result := r.db.First(&spectrum, id); result.Error != nil {
		return result.Error
	}
	spectrum.Image = newImageURL
	result := r.db.Save(spectrum)
	return result.Error
}

func (r *Repository) GetUserRequestID(userID int) (int, error) {
	var userRequestID int
	err := r.db.Table("satellites").Select("id").Where("user_id = ? and status = ?", userID, "черновик").Scan(&userRequestID).Error
	if err != nil {
		return 0, err
	}
	return userRequestID, nil
}
