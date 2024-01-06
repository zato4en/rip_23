package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rip2023/internal/app/repository"
)

// инициализируем структуру репы - логгер и репу

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
}

//Функция, создающая новый хендлер

func NewHandler(l *logrus.Logger, r *repository.Repository) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
	}
}

// эндпоинты

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/", h.SpectrumList)
	router.GET("/Spectrum/:id", h.SpectrumById)
	router.POST("/delete/:id", h.DeleteSpectrum)
	registerStatic(router)
}

// статика
func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}

// ошибки
func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	h.Logger.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}
