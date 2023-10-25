package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"rip2023/internal/app/repository"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Minio      *minio.Client
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.Client) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Minio:      m,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {

	//эндпоинты для услуг
	router.GET("/Spectrums", h.SpectrumsList)
	router.GET("/Spectrumbyid", h.SpectrumById)
	router.POST("/Spectrums", h.AddSpectrum)
	router.PUT("/Spectrums", h.UpdateSpectrum)
	router.DELETE("/Spectrums", h.DeleteSpectrum)

	//эндпоинты для заявок
	router.GET("/Satellites", h.SatellitesList)
	router.DELETE("/Satellites", h.DeleteSatellite)
	router.PUT("/Satellites", h.UpdateSatellite)
	router.PUT("/SatellitesUpdate", h.UpdateSatelliteById)
	router.GET("/UsersSatellite", h.UsersSatellite)

	//эндпоинты для м-м
	//router.GET("/SpectrumsRequests", h.SpectrumRequestsList)
	router.POST("/SpectrumsRequests", h.AddSpectrumToRequest)
	router.DELETE("/SpectrumsRequests", h.DeleteSpectrumRequest)
	router.PUT("/SpectrumsRequests", h.UpdateSpectrumNumberInRequest)

	//эндпоинты для юзеров
	//router.GET("/users", h.UsersList)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}

//request status

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	h.Logger.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}

func (h *Handler) successHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		key:      data,
	})
}

func (h *Handler) successAddHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		key:      data,
	})
}
