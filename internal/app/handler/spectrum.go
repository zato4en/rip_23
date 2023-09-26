package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SpectrumList(ctx *gin.Context) {
	searchQuery := ctx.Query("search")
	if searchQuery == "" {
		Spectrum, err := h.Repository.SpectrumList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Spectrum": Spectrum,
		})
	} else {

		filteredSpectrum, err := h.Repository.SearchSpectrum(searchQuery)
		if err != nil {
			// обработка ошибки
		}
		ctx.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Spectrum": filteredSpectrum,
		})

	}
}

func (h *Handler) SpectrumById(ctx *gin.Context) {
	id := ctx.Param("id")
	Spectrum, err := h.Repository.SpectrumById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "spectrum.html", gin.H{
		"Spectrum": Spectrum,
	})
}

func (h *Handler) DeleteSpectrum(ctx *gin.Context) {
	id := ctx.Param("id")
	h.Repository.DeleteSpectrum(id)
	ctx.Redirect(http.StatusFound, "/")
}
