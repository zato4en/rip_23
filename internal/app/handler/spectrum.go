package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rip2023/internal/app/ds"
)

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
		ctx.JSON(http.StatusOK, gin.H{
			"Spectrums": Spectrums,
		})
	} else {

		filteredSpectrums, err := h.Repository.SearchSpectrum(searchQuery)
		if err != nil {
			// обработка ошибки
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Spectrums": filteredSpectrums,
		})

	}
}

func (h *Handler) SpectrumById(ctx *gin.Context) {
	id := ctx.Param("id")
	Spectrums, err := h.Repository.SpectrumById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Spectrums": Spectrums,
	})
}

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

func (h *Handler) AddSpectrum(ctx *gin.Context) {
	var newSpectrum ds.Spectrum
	if err := ctx.BindJSON(&newSpectrum); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if newSpectrum.ID != 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idMustBeEmpty)
		return
	}
	if newSpectrum.Description == "" {
		h.errorHandler(ctx, http.StatusBadRequest, SpectrumCannotBeEmpty)
		return
	}
	if err := h.Repository.AddSpectrum(&newSpectrum); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	//h.successAddHandler(ctx, "Spectrum_id", newSpectrum.ID)
	ctx.Redirect(http.StatusFound, "/Spectrums")
}

func (h *Handler) UpdateSpectrum(ctx *gin.Context) {
	var updatedSpectrum ds.Spectrum
	if err := ctx.BindJSON(&updatedSpectrum); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedSpectrum.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateSpectrum(&updatedSpectrum); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	h.successHandler(ctx, "updated_Spectrum", gin.H{
		"id":          updatedSpectrum.ID,
		"description": updatedSpectrum.Description,
		"length":      updatedSpectrum.Len,
		"frequency":   updatedSpectrum.Freq,
		"image":       updatedSpectrum.Image,
		"is_delete":   updatedSpectrum.IsDelete,
	})
}
