package utils

import (
	"alqinsidev/auth-service/config"
	"alqinsidev/auth-service/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	Gorm *gorm.DB
}

func NewDBConnection(cfg *config.Config) *Conn {
	return &Conn{Gorm: initGorm(cfg)}
}

func initGorm(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.DB.DSN,
	}), &gorm.Config{})

	if err != nil {
		logrus.Error("Cannot establish POSTGRES connection")
		logrus.Error(cfg.DB.DSN)
	}

	db.AutoMigrate(&domain.User{})
	logrus.Info("MIGRATION OK")

	return db
}
