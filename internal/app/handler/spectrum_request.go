package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rip2023/internal/app/ds"
)

func (h *Handler) SpectrumRequestsList(ctx *gin.Context) {
	SpectrumRequests, err := h.Repository.SpectrumRequestsList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "Spectrums_requests", SpectrumRequests)
}

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
// @Tags Планеты
// @Description Добавление спектра в заявку. Если заявка не найдена, она будет сформирована
// @Accept json
// @Produce json
// @Param request body ds.Spectrum_request true "Данные для добавления спектра в заявку"
// @Success 200 {object} ds.Spectrum_request "ID"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Spectrum_request [post]
func (h *Handler) AddSpectrumToRequest(ctx *gin.Context) {
	//var spectrumRequest ds.spectrumsRequest
	var request struct {
		SpectrumId uint `json:"spectrum_id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	if request.SpectrumId == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, SatelliteIDOrSpectrumIDIsEmpty)
		return
	}

	if err := h.Repository.AddSpectrumToRequest(&request); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successAddHandler(ctx, "updated_spectrum_request", request)
}

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
