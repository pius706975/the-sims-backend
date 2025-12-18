package database

import (
	"fmt"
	envConfig "github.com/pius706975/the-sims-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	envCfg := envConfig.LoadConfig()

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ", envCfg.DBHost, envCfg.DBPort, envCfg.DBUsername, envCfg.DBPassword, envCfg.DBName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDb, nil
}