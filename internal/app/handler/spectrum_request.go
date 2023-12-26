package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rip2023/internal/app/ds"
	"strconv"
)

func (h *Handler) SpectrumRequestsList(ctx *gin.Context) {
	SpectrumRequests, err := h.Repository.SpectrumRequestsList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "Spectrums_requests", SpectrumRequests)
}

// UpdateSpectrumNumberInRequest godoc
// @Summary Обновление порядка исследования спектров
// @Description Обновление порядкого номера спектра в заявке
// @Tags Спектры в заявках
// @Accept json
// @Produce json
// @Param request body ds.UpdateSpectrumInRequestNumberReq true "Данные для добавления спектра в заявку"
// @Success 200 {object} ds.UpdateSpectrumInRequestNumberRes "Updated Spectrum In Request ID"
// @Failure 400 {object} errorResp "Плохой запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибку"
// @Router /SpectrumsRequests [put]
func (h *Handler) UpdateSpectrumNumberInRequest(ctx *gin.Context) {
	var updatedSpectrumRequest ds.Spectrum_request
	if err := ctx.BindJSON(&updatedSpectrumRequest); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedSpectrumRequest.SpectrumID == 0 || updatedSpectrumRequest.SatelliteID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateSpectrumNumberInRequest(&updatedSpectrumRequest); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	h.successHandler(ctx, "updated_Spectrum", gin.H{
		"Satellite_id":     updatedSpectrumRequest.SatelliteID,
		"Spectrum_id":      updatedSpectrumRequest.SpectrumID,
		"Satellite_number": updatedSpectrumRequest.Satellite_number,
	})
}

// AddSpectrumToRequest godoc
// @Summary Добавление спектра в заявку
// @Security ApiKeyAuth
// @Tags Спектры в заявках
// @Description Добавление спектра в заявку. Если заявка не найдена, она будет сформирована
// @Accept json
// @Produce json
// @Param request body ds.AddSpectrumToRequestReq true "Данные для добавления планеты в заявку"
// @Success 200 {object} ds.AddSpectrumToRequestResp "ID"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /SpectrumsRequests [post]
func (h *Handler) AddSpectrumToRequest(ctx *gin.Context) {
	// Получение значения userid из контекста
	userID, exists := ctx.Get("user_id")
	if !exists {
		// Обработка ситуации, когда userid отсутствует в контексте
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id not found in context"))
		return
	}
	// Приведение типа, если необходимо
	var userIDUint uint
	switch v := userID.(type) {
	case uint:
		userIDUint = v
	case int:
		userIDUint = uint(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errors.New("failed to convert user_id to uint"))
			return
		}
		userIDUint = uint(i)
	default:
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id is not of a supported type"))
		return
	}
	//var SpectrumRequest ds.SpectrumsRequest
	var request struct {
		SpectrumId uint `json:"Spectrum_id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.SpectrumId == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, SatelliteIDOrSpectrumIDIsEmpty)
		return
	}

	if err := h.Repository.AddSpectrumToRequest(&request, userIDUint); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successAddHandler(ctx, "updated_Spectrum_request", request)
}

// DeleteSpectrumRequest godoc
// @Summary Удаление спектра из заявки
// @Description Удаление спектра из заявки по идентификатору
// @Tags Спектры в заявках
// @Accept json
// @Produce json
// @Param request body ds.DeleteSpectrumInRequestReq true "Идентификатор спектра в заявке"
// @Success 200 {object} ds.DeleteSpectrumInRequestRes "Удаленный идентификатор спектра"
// @Failure 400 {object} errorResp "Плохой запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /SpectrumsRequests [delete]
func (h *Handler) DeleteSpectrumRequest(ctx *gin.Context) {
	var request struct {
		SatelliteID uint `json:"satellite_id"`
		SpectrumID  uint `json:"spectrum_id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.SatelliteID == 0 || request.SpectrumID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeleteSpectrumRequest(request.SatelliteID, request.SpectrumID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "satellite_id", request.SatelliteID)
}
