package repository

import (
	"rip2023/internal/app/ds"
	"strconv"
	"strings"
)

// формирование списка спектров
func (r *Repository) SpectrumList() (*[]ds.Spectrum, error) {
	var Spectrum []ds.Spectrum
	r.db.Where("is_delete = ?", false).Find(&Spectrum)
	return &Spectrum, nil
}

// поиск по спектрам
func (r *Repository) SearchSpectrum(search string) (*[]ds.Spectrum, error) {
	var Spectrum []ds.Spectrum
	r.db.Find(&Spectrum)

	var filteredSpectrum []ds.Spectrum
	for _, Spectrum := range Spectrum {
		if strings.Contains(strings.ToLower(Spectrum.Description), strings.ToLower(search)) {
			filteredSpectrum = append(filteredSpectrum, Spectrum)
		}
	}

	return &filteredSpectrum, nil
}

// спектр по id
func (r *Repository) SpectrumById(id string) (*ds.Spectrum, error) {
	var Spectrum ds.Spectrum
	intId, _ := strconv.Atoi(id)
	r.db.Find(&Spectrum, intId)
	return &Spectrum, nil
}

// удаление спектра (установка флажка isDelete = true)
func (r *Repository) DeleteSpectrum(id string) {
	query := "UPDATE Spectrums SET is_delete = true WHERE id = $1"
	r.db.Exec(query, id)
}
