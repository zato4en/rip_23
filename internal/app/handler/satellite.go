package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rip2023/internal/app/ds"
	"strconv"
)

func (h *Handler) SatellitesList(ctx *gin.Context) {
	Satellites, err := h.Repository.SatellitesList()
	if SatelliteIdString := ctx.Query("Satellites"); SatelliteIdString != "" {
		SatelliteById(ctx, h, SatelliteIdString)
		return
	}

	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Satellites", Satellites)
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
	//ctx.Redirect(http.StatusOK, "/Satellites")
	//h.SatellitesList(ctx)
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
