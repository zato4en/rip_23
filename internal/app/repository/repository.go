package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// инициализируем структуру репы - логгер и БД

type Repository struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewRepository(dsn string, log *logrus.Logger) (*Repository, error) {
	//метод создания репозитория

	//инициализируем GORM
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//возвращаем созданную репу
	return &Repository{
		db:     gormDB,
		logger: log,
	}, nil
}
