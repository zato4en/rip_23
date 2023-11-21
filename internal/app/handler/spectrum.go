package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"rip2023/internal/app/ds"
	"strconv"
)

// /Users/drakenchef/go/bin/swag init -g cmd/main/main.go

// SpectrumsList godoc
// @Summary Список спектров
// @Description Получение спектров и фильтрация при поиске
// @Tags Спектры
// @Produce json
// @Param Spectrum query string false "Получаем определённый спектр"
// @Param search query string false "Фильтрация поиска"
// @Success 200 {object} ds.SpectrumsListResp
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Spectrums [get]
func (h *Handler) SpectrumsList(ctx *gin.Context) {
	searchQuery := ctx.Query("search")
	if searchQuery == "" {
		Spectrums, err := h.Repository.SpectrumsList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Получаем id заявки пользователя
		userRequestID, err := h.Repository.GetUserRequestID(1)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Spectrums": Spectrums,
			"Flight_id": userRequestID,
		})
	} else {
		filteredSpectrums, err := h.Repository.SearchSpectrum(searchQuery)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Получаем id заявки пользователя
		userRequestID, err := h.Repository.GetUserRequestID(1)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Spectrums": filteredSpectrums,
			"Flight_id": userRequestID,
		})
	}
}

func (h *Handler) SpectrumById(ctx *gin.Context) {
	id := ctx.Param("id")
	idint, _ := strconv.Atoi(id)
	spectrums, err := h.Repository.SpectrumById(idint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Spectrums": spectrums,
	})
}

// DeleteSpectrum godoc
// @Summary Удаление Спектра
// @Description Удаление Спектра по его идентификатору.
// @Security ApiKeyAuth
// @Tags Спектры
// @Accept json
// @Produce json
// @Param request body ds.DeleteSpectrumReq true "ID Спектра для удаления"
// @Success 200 {object} ds.DeleteSpectrumRes "Спектр успешно удален"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Spectrums [delete]
func (h *Handler) DeleteSpectrum(ctx *gin.Context) {
	//id := ctx.Param("id")
	//h.Repository.DeleteSpectrum(id)
	//ctx.Redirect(http.StatusFound, "/Spectrums")
	var request struct {
		ID uint `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeleteSpectrum(request.ID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "Spectrum_id", request.ID)
}

// AddSpectrum godoc
// @Summary Создание Спектра
// @Security ApiKeyAuth
// @Tags Спектры
// @Description Создание Спектра
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "Название Спектра"
// @Param status formData string true "Статус Спектра"
// @Param description formData string true "Описание Спектра"
// @Param image formData file true "Изображение Спектра"
// @Success 201 {object} ds.AddSpectrumResp
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Spectrums [post]
func (h *Handler) AddSpectrum(ctx *gin.Context) {
	length := ctx.Request.FormValue("len")
	freq := ctx.Request.FormValue("freq")
	description := ctx.Request.FormValue("description")
	lengthfloat, _ := strconv.ParseFloat(length, 64)
	freqfloat, _ := strconv.ParseFloat(freq, 64)

	newSpectrum := ds.Spectrum{
		IsDelete:    false,
		Description: description,
		Len:         lengthfloat,
		Freq:        freqfloat,
	}
	file, header, _ := ctx.Request.FormFile("image")
	if errorCode, errCreate := h.createSpectrum(&newSpectrum); errCreate != nil {
		h.errorHandler(ctx, errorCode, errCreate)
	}
	if file != nil && header.Size != 0 && header != nil {
		newImageURL, errCode, errDB := h.createImageSpectrum(&file, header, fmt.Sprintf("%d", newSpectrum.ID))
		if errDB != nil {
			h.errorHandler(ctx, errCode, errDB)
			return
		}
		newSpectrum.Image = newImageURL
	}
	ctx.Redirect(http.StatusFound, "/Spectrums")
}

func (h *Handler) createSpectrum(Spectrum *ds.Spectrum) (int, error) {
	if Spectrum.ID != 0 {
		return http.StatusBadRequest, idMustBeEmpty
	}
	if Spectrum.Description == "" {
		return http.StatusBadRequest, SpectrumCannotBeEmpty
	}
	if err := h.Repository.AddSpectrum(Spectrum); err != nil {
		return http.StatusBadRequest, err
	}
	return 0, nil
}

// UpdateSpectrum godoc
// @Summary Обновление информации о спектре
// @Security ApiKeyAuth
// @Tags Спектры
// @Description Обновление информации о спектре
// @Accept json
// @Produce json
// @Param updated_Spectrum body ds.UpdateSpectrumReq true "Обновленная информация о спектре"
// @Success 200 {object} ds.UpdateSpectrumResp
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Spectrums [put]
func (h *Handler) UpdateSpectrum(ctx *gin.Context) {

	spectrumId := ctx.Request.FormValue("id")
	length := ctx.Request.FormValue("len")
	freq := ctx.Request.FormValue("freq")
	description := ctx.Request.FormValue("description")
	lengthfloat, _ := strconv.ParseFloat(length, 64)
	freqfloat, _ := strconv.ParseFloat(freq, 64)
	spectrumIduint, _ := strconv.Atoi(spectrumId)

	newSpectrum := ds.Spectrum{
		ID:          uint(spectrumIduint),
		IsDelete:    false,
		Description: description,
		Len:         lengthfloat,
		Freq:        freqfloat,
	}
	file, header, _ := ctx.Request.FormFile("image")
	if errorCode, errCreate := h.updateSpectrum(&newSpectrum); errCreate != nil {
		h.errorHandler(ctx, errorCode, errCreate)
	}
	if file != nil && header.Size != 0 && header != nil {
		newImageURL, errCode, errDB := h.createImageSpectrum(&file, header, fmt.Sprintf("%d", newSpectrum.ID))
		if errDB != nil {
			h.errorHandler(ctx, errCode, errDB)
			return
		}
		newSpectrum.Image = newImageURL
	}
	ctx.Redirect(http.StatusFound, "/Spectrums")

}

// asd
func (h *Handler) updateSpectrum(Spectrum *ds.Spectrum) (int, error) {
	if Spectrum.ID == 0 {
		return http.StatusBadRequest, idNotFound
	}
	if err := h.Repository.UpdateSpectrum(Spectrum); err != nil {
		return http.StatusBadRequest, err
	}
	return 0, nil
}

func (h *Handler) createImageSpectrum(
	file *multipart.File,
	header *multipart.FileHeader,
	SpectrumID string,
) (string, int, error) {
	newImageURL, errMinio := h.createImageInMinio(file, header)
	if errMinio != nil {
		return "", http.StatusInternalServerError, errMinio
	}
	if err := h.Repository.UpdateSpectrumImage(SpectrumID, newImageURL); err != nil {
		return "", http.StatusInternalServerError, err
	}
	return newImageURL, 0, nil
}
