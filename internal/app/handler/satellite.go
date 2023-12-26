package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"rip2023/internal/app/ds"
	"strconv"
	"time"
)

// SatellitesList godoc
// @Summary Список заявок
// @Tags Заявки
// @Security ApiKeyAuth
// @Description Получение списка заявок с фильтрами по статусу, дате начала и дате окончания, пользователю.
// @Produce json
// @Param status query string false "Статус заявки."
// @Param date_formation_start query string false "Дата начала периода фильтрации в формате '2006-01-02'."
// @Param date_formation_end query string false "Дата окончания периода фильтрации в формате '2006-01-02'."
// @Param user_id query int false "ID пользователя"
// @Success 200 {array} ds.SatellitesListRes "Список заявок"
// @Success 200 {array} ds.SatellitesListRes2 "Список заявок"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 204 {object} errorResp "Нет данных"
// @Router /Satellites [get]
func (h *Handler) SatellitesList(ctx *gin.Context) {
	// Получение значения userid из контекста
	userID, exists := ctx.Get("user_id")
	if !exists {
		// Обработка ситуации, когда userid отсутствует в контексте
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id not found in context"))
		return
	}

	// Приведение типа, если необходимо
	var userIDUint uint
	switch v := userID.(type) {
	case uint:
		userIDUint = v
	case int:
		userIDUint = uint(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errors.New("failed to convert user_id to uint"))
			return
		}
		userIDUint = uint(i)
	default:
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id is not of a supported type"))
		return
	}
	var thisUser *ds.Users
	thisUser = h.Repository.GetUserById(userIDUint)
	if thisUser.Role == 0 {
		Satellite, err := h.Repository.UsersSatellite(userIDUint)
		if err != nil {
			h.errorHandler(ctx, http.StatusNoContent, err)
			return
		}
		h.successHandler(ctx, "Satellites", Satellite)

	} else {
		userlogin := ctx.DefaultQuery("user_login", "")
		datestart := ctx.DefaultQuery("date_formation_start", "")
		dateend := ctx.DefaultQuery("date_formation_end", "")
		status := ctx.DefaultQuery("status", "")

		Satellites, err := h.Repository.SatellitesList(userlogin, datestart, dateend, status)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Satellites"})
			return
		}

		type SatelliteResponse struct {
			ID             uint      `json:"id"`
			DateCreate     time.Time `json:"date_create"`
			DateFormation  time.Time `json:"date_formation"`
			DateCompletion time.Time `json:"date_completion"`
			Status         string    `json:"status"`
			Satellite      string    `json:"satellite"`
			ModerLogin     string    `json:"moder_login"`
			UserLogin      string    `json:"user_login"`
			Percentage     string    `json:"percentage"`
		}

		SatelliteResponses := []SatelliteResponse{}
		for _, Satellite := range *Satellites {
			SatelliteResponse := SatelliteResponse{
				ID:             Satellite.ID,
				DateCreate:     Satellite.DateCreate,
				DateFormation:  Satellite.DateFormation,
				DateCompletion: Satellite.DateCompletion,
				Status:         Satellite.Status,
				ModerLogin:     Satellite.ModerLogin,
				UserLogin:      Satellite.UserLogin,
				Percentage:     Satellite.Percentage,
			}
			SatelliteResponses = append(SatelliteResponses, SatelliteResponse)
		}

		// Отправка измененного JSON-ответа без user_id и moder_id
		//ctx.JSON(http.StatusOK, SatelliteResponses)
		h.successHandler(ctx, "Satellites", SatelliteResponses)

	}
}

// UsersSatellite godoc
// @Summary Список заявок пользователя
// @Tags Заявки
// @Security ApiKeyAuth
// @Description Получение списка заявок пользователем.
// @Produce json
// @Success 200 {array} ds.SatellitesListRes "Список заявок"
// @Success 200 {array} ds.SatellitesListRes2 "Список заявок"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 204 {object} errorResp "Нет данных"
// @Router /UsersSatellite [get]
func (h *Handler) UsersSatellite(ctx *gin.Context) {
	// Получение значения userid из контекста
	userID, exists := ctx.Get("user_id")
	if !exists {
		// Обработка ситуации, когда userid отсутствует в контексте
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id not found in context"))
		return
	}

	// Приведение типа, если необходимо
	var userIDUint uint
	switch v := userID.(type) {
	case uint:
		userIDUint = v
	case int:
		userIDUint = uint(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errors.New("failed to convert user_id to uint"))
			return
		}
		userIDUint = uint(i)
	default:
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id is not of a supported type"))
		return
	}

	Satellite, err := h.Repository.UsersSatellite(userIDUint)
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Satellite", Satellite)
}

// DeleteSatellite godoc
// @Summary Удаление заявки
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Удаление заявки по идентификатору.
// @Accept json
// @Produce json
// @Param request body ds.DeleteSatelliteReq true "Идентификатор заявки для удаления"
// @Success 200 {object} ds.DeleteSatelliteRes "Успешное удаление заявки"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Satellites [delete]
func (h *Handler) DeleteSatellite(ctx *gin.Context) {
	var request struct {
		ID uint `json:"id"`
	}
	//
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

// Импортируем пакет errors

func (h *Handler) UpdateSatelliteAsyncStatus(ctx *gin.Context) {
	var req struct {
		Percentage string `json:"percentage"`
		Passkey    string `json:"passkey"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("error parsing request body"))
		return
	}

	if req.Passkey != "password" {
		h.errorHandler(ctx, http.StatusForbidden, errors.New("invalid passkey"))
		return
	}

	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("id not found"))
		return
	}
	if len(req.Percentage) == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, errors.New("percentage not found"))
		return
	}
	//koment
	if err := h.Repository.UpdateSatelliteAsyncStatus(idint, req.Percentage); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("error updating status"))
		return
	}

	h.successHandler(ctx, "percentage_updated", gin.H{
		"satellite_id": idint,
		"Percentage":   req.Percentage,
	})
}

// UpdateSatellite godoc
// @Summary Обновление данных о заявке
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление данных о заявке.
// @Accept json
// @Produce json
// @Param updatedSatellite body ds.UpdateSatelliteReq true "Данные для обновления заявки"
// @Success 200 {object} ds.UpdatedSatelliteRes "Успешное обновление данных о заявке"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /Satellites [put]
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
		"date_created":  updatedSatellite.DateCreate,
		"date_formed":   updatedSatellite.DateFormation,
		"date_accepted": updatedSatellite.DateCompletion,
		"status":        updatedSatellite.Status,
		"satellite":     updatedSatellite.Satellite,
		"user_id":       updatedSatellite.UserID,
		"moder_id":      updatedSatellite.ModerID,
	})
}

// UsersUpdateSatellite godoc
// @Summary Обновление данных о заявке пользователем
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление данных о заявке пользователем.
// @Accept json
// @Produce json
// @Param updatedHike body ds.UpdateSatelliteReq true "Данные для обновления заявки пользователем"
// @Success 200 {object} ds.UpdatedSatelliteRes "Успешное обновление данных о заявке пользователя"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /UsersSatelliteUpdate [put]
func (h *Handler) UsersUpdateSatellite(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		// Обработка ситуации, когда userid отсутствует в контексте
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id not found in context"))
		return
	}
	// Приведение типа, если необходимо
	var userIDUint uint
	switch v := userID.(type) {
	case uint:
		userIDUint = v
	case int:
		userIDUint = uint(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errors.New("failed to convert user_id to uint"))
			return
		}
		userIDUint = uint(i)
	default:
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("user_id is not of a supported type"))
		return
	}

	var updatedSatellite ds.Satellite
	if err := ctx.BindJSON(&updatedSatellite); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	//if updatedSatellite.ID == 0 {
	//	h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
	//	return
	//}
	if err := h.Repository.UsersUpdateSatellite(&updatedSatellite, userIDUint); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "updated_Satellite", gin.H{
		"id":            updatedSatellite.ID,
		"date_created":  updatedSatellite.DateCreate,
		"date_formed":   updatedSatellite.DateFormation,
		"date_accepted": updatedSatellite.DateCompletion,
		"status":        updatedSatellite.Status,
		"satellite":     updatedSatellite.Satellite,
		"user_id":       updatedSatellite.UserID,
		"moder_id":      updatedSatellite.ModerID,
	})
}

// UserUpdateSatelliteStatusById godoc
// @Summary Обновление статуса заявки для пользователя.
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление статуса заявки для пользователя.
// @Accept json
// @Produce json
// @Param id path string true "ID заявки"
// @Success 200 {object} string "Успешное обновление статуса"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /SatellitesUser/{id} [put]
func (h *Handler) UserUpdateSatelliteStatusById(ctx *gin.Context) {
	id := ctx.Param("id")

	// Создаем структуру для запроса
	requestBody, err := json.Marshal(map[string]string{
		"satellite_id": id,
	})
	if err != nil {
		// Обработка ошибки маршалинга JSON
		ctx.String(http.StatusInternalServerError, "Error creating request body: %v", err)
		return
	}

	// Отправляем запрос на внешний сервис
	resp, err := http.Post("http://localhost:8000/start-async-update/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		// Обработка ошибки выполнения запроса
		ctx.String(http.StatusInternalServerError, "Error sending request to the external service: %v", err)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		// Обработка случая, когда внешний сервис вернул ошибку
		ctx.String(resp.StatusCode, "External service returned: %s", resp.Status)
		return
	}

	// Все хорошо, возвращаем HTTP статус 200 OK
	ctx.Status(http.StatusOK)
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

// ModerUpdateSatelliteStatusById godoc
// @Summary Обновление статуса заявки для модератора
// @Security ApiKeyAuth
// @Tags Заявки
// @Description Обновление статуса заявки для модератора.
// @Accept json
// @Produce json
// @Param id path string true "ID заявки"
// @Success 200 {object} string "Успешное обновление статуса"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 500 {object} errorResp "Внутренняя ошибка сервера"
// @Router /SatellitesModer/{id} [put]
func (h *Handler) ModerUpdateSatelliteStatusById(ctx *gin.Context) {
	var requestData struct {
		Status    string `json:"status"`
		Modername string `json:"modername"`
	}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	status := requestData.Status
	modername := requestData.Modername

	id := ctx.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	result, err := h.Repository.ModerUpdateSatelliteStatusById(idint, modername, status)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errors.New("can not refactor status"))
		return
	}

	h.successHandler(ctx, "updated_status_by_moder", gin.H{
		"id":     result.ID,
		"status": result.Status,
	})
}

// SatelliteById godoc
// @Summary Получение информации о заявке по её ID.
// @Tags Заявки
// @Description Получение информации о заявке по его ID.
// @Produce json
// @Param id path string true "ID заявки"
// @Success 200 {object} ds.SatellitesListRes2 "Информация о заявке по ID"
// @Failure 400 {object} errorResp "Неверный запрос"
// @Failure 404 {object} errorResp "Заявка не найдена"
// @Router /Satellites/{id} [get]
func (h *Handler) SatelliteById(ctx *gin.Context) {
	id := ctx.Param("id")
	Satellite, err := h.Repository.SatelliteById(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNoContent, err)
		return
	}
	h.successHandler(ctx, "Satellite", Satellite)

}
