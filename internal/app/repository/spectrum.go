package repository

import (
	"rip2023/internal/app/ds"
	"strconv"
	"strings"
)

//Этот файл - описание методов сущности (услуг) на уровне репы

// формирование списка спектров

func (r *Repository) SpectrumList() (*[]ds.Spectrum, error) {
	var Spectrum []ds.Spectrum
	//Поиск с условием, используя методы ГОРМ
	r.db.Where("is_delete = ?", false).Find(&Spectrum)
	return &Spectrum, nil
}

// поиск по спектрам
//ПЕРЕДЕЛАЛ НА ОРМ

func (r *Repository) SearchSpectrum(search string) (*[]ds.Spectrum, error) {
	var filteredSpectrum []ds.Spectrum
	search = "%" + strings.ToLower(search) + "%"
	if err := r.db.Where("LOWER(description) LIKE ?", search).Find(&filteredSpectrum).Error; err != nil {
		return nil, err
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
	//В данном случае мы пробуем сделать это напрямую SQL запросом, а не ГОРМом
	query := "UPDATE Spectrums SET is_delete = true WHERE id = $1"
	r.db.Exec(query, id)
}
