package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rip2023/MyMinio"
	"rip2023/internal/app/config"
	"rip2023/internal/app/dsn"
	"rip2023/internal/app/handler"
	app "rip2023/internal/app/pkg"
	"rip2023/internal/app/redis"
	"rip2023/internal/app/repository"
)

// @title Spectrum Analysis
// @version 1.0
// @description Here is the API for CMB spectrum analylis requests.
// @contact.name API Support
// @contact.url https://github.com/zato4en
// @contact.email lavrenovmihail2103@gmail.com

// @host localhost:8888
// @schemes http
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// ShowAccount godoc
// @Summary      Spectrums
// @Description  Get spectrums list
// @Tags         spectrums
// @Produce      json
// @Router       /Spectrums [get]

func main() {
	logger := logrus.New()
	minioClient := MyMinio.NewMinioClient(logger)
	router := gin.Default()
	conf, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Error with configuration reading: %s", err)
	}
	ctx := context.Background()
	redisClient, errRedis := redis.New(ctx, conf.Redis)
	if errRedis != nil {
		logger.Fatalf("Errof with redis connect: %s", err)
	}
	postgresString, errPost := dsn.FromEnv()
	if errPost != nil {
		logger.Fatalf("Error of reading postgres line: %s", errPost)
	}
	fmt.Println(postgresString)
	rep, errRep := repository.NewRepository(postgresString, logger)
	if errRep != nil {
		logger.Fatalf("Error from repository: %s", err)
	}
	hand := handler.NewHandler(logger, rep, minioClient, conf, redisClient)
	application := app.NewApp(conf, router, logger, hand)
	application.RunApp()
}
