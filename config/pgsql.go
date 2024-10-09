package config

import (
	"ecommerce/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg Config) {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	DB, err = gorm.Open(postgres.Open(psqlInfo))
	if err != nil {
		panic("failed to connect to the database")
	}

	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Order{})
}
