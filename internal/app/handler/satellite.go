package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rip2023/internal/app/ds"
	"strconv"
)

func (h *Handler) SatellitesList(ctx *gin.Context) {
	userID := ctx.DefaultQuery("user_id", "")
	datestart := ctx.DefaultQuery("date_start", "")
	dateend := ctx.DefaultQuery("date_end", "")
	status := ctx.DefaultQuery("status", "")

	Satellites, err := h.Repository.SatellitesList(userID, datestart, dateend, status)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch satellites"})
		return
	}

	ctx.JSON(http.StatusOK, Satellites)
}

func (h *Handler) SatelliteById(ctx *gin.Context) {
	id := ctx.Param("id")
	Satellite, err := h.Repository.SatelliteById(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Satellite", Satellite)

}

func (h *Handler) UserUpdateSatelliteStatusById(ctx *gin.Context) {
	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	result, err := h.Repository.UserUpdateSatelliteStatusById(idint)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("can not refactor status"))
		return
	}

	h.successHandler(ctx, "updated_status_by_user", gin.H{
		"id":     result.ID,
		"status": result.Status,
	})
}

func (h *Handler) ModerUpdateSatelliteStatusById(ctx *gin.Context) {
	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	result, err := h.Repository.ModerUpdateSatelliteStatusById(idint)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("can not refactor status"))
		return
	}

	h.successHandler(ctx, "updated_status_by_moder", gin.H{
		"id":     result.ID,
		"status": result.Status,
	})
}

func (h *Handler) UpdateSatelliteById(ctx *gin.Context) {
	var updatedSatellite ds.Satellite
	if err := ctx.BindJSON(&updatedSatellite); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedSatellite.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateSatelliteStatus(&updatedSatellite); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "updated_satellite", gin.H{
		"id":     updatedSatellite.ID,
		"status": updatedSatellite.Status,
	})
}

func (h *Handler) DeleteSatellite(ctx *gin.Context) {
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
	if err := h.Repository.DeleteSatellite(request.ID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "Satellite_id", request.ID)

}

func (h *Handler) UpdateSatellite(ctx *gin.Context) {
	var updatedSatellite ds.Satellite
	if err := ctx.BindJSON(&updatedSatellite); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedSatellite.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateSatellite(&updatedSatellite); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}
	h.successHandler(ctx, "updated_Satellite", gin.H{
		"id":            updatedSatellite.ID,
		"date_created":  updatedSatellite.DateCreated,
		"date_formed":   updatedSatellite.DateFormed,
		"date_accepted": updatedSatellite.DateAccepted,
		"status":        updatedSatellite.Status,
		"satellite":     updatedSatellite.Satellite,
		"user_id":       updatedSatellite.UserID,
		"moder_id":      updatedSatellite.ModerID,
	})

}

func (h *Handler) UsersSatellite(ctx *gin.Context) {
	Satellite, err := h.Repository.UsersSatellite()
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Satellite", Satellite)
}
