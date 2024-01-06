package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rip2023/internal/app/config"
	"rip2023/internal/app/handler"
)

//Создаем структуру приложения

type Application struct {
	Config  *config.Config
	Logger  *logrus.Logger
	Router  *gin.Engine
	Handler *handler.Handler
}

//Функция создающая новую структуру

func NewApp(c *config.Config, r *gin.Engine, l *logrus.Logger, h *handler.Handler) *Application {
	return &Application{
		Config:  c,
		Logger:  l,
		Router:  r,
		Handler: h,
	}
}

//Функция запуска приложения

func (a *Application) RunApp() {
	a.Logger.Info("Server start up")
	//Обращаемся к полю структуры Application - хендлеру
	//Используем ранее написанный в нем метод RegisterHandler
	//Передаем в него роутер
	a.Handler.RegisterHandler(a.Router)

	serverAddress := fmt.Sprintf("%s:%d", a.Config.ServiceHost, a.Config.ServicePort)
	if err := a.Router.Run(serverAddress); err != nil {
		a.Logger.Fatalln(err)
	}
	a.Logger.Info("Server down")
}
