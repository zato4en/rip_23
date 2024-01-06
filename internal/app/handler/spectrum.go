package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Этот файл - описание методов сущности (услуг) на уровне хендлера

//Метод вывода всех услуг

func (h *Handler) SpectrumList(ctx *gin.Context) {
	spectrumName := ctx.Query("search")
	if spectrumName == "" {
		Spectrum, err := h.Repository.SpectrumList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Spectrum": Spectrum,
			//"SearchName": spectrumName,
		})
	} else {

		filteredSpectrum, err := h.Repository.SearchSpectrum(spectrumName)
		if err != nil {
			// обработка ошибки
		}
		ctx.HTML(http.StatusOK, "mainpage.html", gin.H{
			"Spectrum": filteredSpectrum,
		})

	}
}

//Метод

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
