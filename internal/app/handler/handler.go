package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"rip2023/internal/app/config"
	_ "rip2023/internal/app/docs"
	redis2 "rip2023/internal/app/redis"
	"rip2023/internal/app/repository"
	"rip2023/internal/app/role"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Minio      *minio.Client
	Config     *config.Config
	Redis      *redis2.Client
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.Client, conf *config.Config, red *redis2.Client) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Minio:      m,
		Config:     conf,
		Redis:      red,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.UserCRUD(router)
	h.SpectrumCRUD(router)
	h.SatelliteCRUD(router)
	h.SpectrumsRequestsCRUD(router)
	registerStatic(router)
}
func (h *Handler) UserCRUD(router *gin.Engine) {
	router.GET("/users", h.UsersList)
	router.POST("/login", h.Login)
	router.POST("/signup", h.Register)
	router.GET("/logout", h.Logout)
}
func (h *Handler) SpectrumCRUD(router *gin.Engine) {
	router.GET("/api/Spectrums", h.WithoutAuthCheck(role.Buyer, role.Manager, role.Admin), h.SpectrumsList)
	router.GET("/api/Spectrums/:id", h.SpectrumById)
	router.POST("/api/Spectrums", h.WithAuthCheck(role.Manager, role.Admin), h.AddSpectrum)
	router.PUT("/api/Spectrums/:id", h.WithAuthCheck(role.Manager, role.Admin), h.UpdateSpectrum)
	router.DELETE("/api/Spectrums", h.WithAuthCheck(role.Manager, role.Admin), h.DeleteSpectrum)
}
func (h *Handler) SatelliteCRUD(router *gin.Engine) {
	router.GET("/api/Satellites", h.WithAuthCheck(role.Manager, role.Admin), h.SatellitesList)
	router.GET("/api/Satellites/:id", h.WithAuthCheck(role.Manager, role.Admin), h.SatelliteById)
	router.DELETE("/api/Satellites", h.WithAuthCheck(role.Manager, role.Admin), h.DeleteSatellite)
	router.PUT("/api/Satellites", h.WithIdCheck(role.Manager, role.Admin), h.UpdateSatellite)
	router.PUT("/api/SatellitesAsyncStatus/:id", h.UpdateSatelliteAsyncStatus)
	router.PUT("/api/SatellitesUser/:id", h.WithAuthCheck(role.Buyer), h.UserUpdateSatelliteStatusById)
	router.PUT("/api/SatellitesModer/:id", h.WithAuthCheck(role.Manager, role.Admin), h.ModerUpdateSatelliteStatusById)
	router.GET("/api/UsersSatellite", h.WithIdCheck(role.Buyer, role.Manager, role.Admin), h.UsersSatellite)
	router.PUT("/api/UsersSatelliteUpdate", h.WithIdCheck(role.Buyer, role.Manager, role.Admin), h.UsersUpdateSatellite)
}
func (h *Handler) SpectrumsRequestsCRUD(router *gin.Engine) {
	router.POST("/api/SpectrumsRequests", h.WithIdCheck(role.Buyer, role.Manager, role.Admin), h.AddSpectrumToRequest)
	router.DELETE("/api/SpectrumsRequests", h.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.DeleteSpectrumRequest)
	router.PUT("/api/SpectrumsRequests", h.WithAuthCheck(role.Buyer, role.Manager, role.Admin), h.UpdateSpectrumNumberInRequest)
	router.GET("/api/ping", h.WithAuthCheck(role.Manager, role.Admin), h.Ping)
}

func registerStatic(router *gin.Engine) {
	router.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/static", "./static")
	router.Static("/img", "./static")
}

// request status

// MARK: - Error handler
type errorResp struct {
	Status      string `json:"status" example:"error"`
	Description string `json:"description" example:"Описание ошибки"`
}

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

// Ping godoc
// @Summary       hello text
// @Description
// @Tags         Тестик
// @Security ApiKeyAuth
// @Produce      json
// @Router       /ping [get]
func (h *Handler) Ping(gCtx *gin.Context) {
	name := gCtx.Request.FormValue("name")
	gCtx.String(http.StatusOK, "Hello, %s", name)
}
