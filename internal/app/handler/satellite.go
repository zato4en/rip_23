package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rip2023/internal/app/ds"
	"strconv"
)

func (h *Handler) SatellitesList(ctx *gin.Context) {
	Satellites, err := h.Repository.SatellitesList()

	//user id sort in requests
	if userIdString := ctx.Query("Sort"); userIdString == "ID" {
		var request struct {
			ID uint `json:"id"`
		}
		if err = ctx.BindJSON(&request); err != nil {
			h.errorHandler(ctx, http.StatusBadRequest, err)
			return
		}
		if request.ID == 0 {
			h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
			return
		}
		var Satellite *[]ds.Satellite
		if Satellite, err = h.Repository.SatellitesListByUser(request.ID); err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
			return
		}
		h.successHandler(ctx, "Satellites by user id", Satellite)
		return
	}

	//date sort in requests
	if DateString := ctx.Query("Sort"); DateString == "Date" {
		var request struct {
			DateFormationStart string `json:"date_formation_start" time_format:"2006-01-02"`
			DateFormationEnd   string `json:"date_formation_end" time_format:"2006-01-02"`
		}
		if err = ctx.BindJSON(&request); err != nil {
			h.errorHandler(ctx, http.StatusBadRequest, err)
			return
		}
		if request.DateFormationStart == "" {
			h.errorHandler(ctx, http.StatusBadRequest, errors.New("empty date input"))
			return
		}
		if request.DateFormationEnd == "" {
			h.errorHandler(ctx, http.StatusBadRequest, errors.New("empty date input"))
			return
		}
		var Satellite *[]ds.Satellite
		if Satellite, err = h.Repository.SatellitesListByDate(request.DateFormationStart, request.DateFormationEnd); err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
			return
		}
		h.successHandler(ctx, "Satellites by date", Satellite)
		return
	}

	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Satellites", Satellites)
}

func (h *Handler) UsersSatellite(ctx *gin.Context) {
	satellite, err := h.Repository.UsersSatellite()
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Satellite", satellite)
}

func SatelliteById(ctx *gin.Context, h *Handler, SatelliteStringID string) {
	SatelliteID, err := strconv.Atoi(SatelliteStringID)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	Satellite, errDB := h.Repository.SatelliteById(uint(SatelliteID))
	if errDB != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errDB)
		return
	}
	h.successHandler(ctx, "Satellite", Satellite)
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
