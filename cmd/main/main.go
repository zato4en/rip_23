package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rip2023/internal/app/config"
	"rip2023/internal/app/dsn"
	"rip2023/internal/app/handler"
	"rip2023/internal/app/pkg"
	"rip2023/internal/app/repository"
)

func main() {
	logger := logrus.New()
	router := gin.Default()
	conf, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Error with configuration reading: %s", err)
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
	hand := handler.NewHandler(logger, rep)
	application := pkg.NewApp(conf, router, logger, hand)
	application.RunApp()
}
