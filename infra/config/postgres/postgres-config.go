package postgres

import (
	"fmt"
	"go-api/infra/config/env"
	"go-api/infra/repositories/postgres/models"

	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := env.GetEnv("DB_HOST", "localhost")
	dbPort := env.GetEnv("DB_PORT", "5432")
	dbUser := env.GetEnv("DB_USER", "postgres")
	dbPassword := env.GetEnv("DB_PASSWORD", "postgres")
	dbName := env.GetEnv("DB_NAME", "go_api_db")

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	database, err := gorm.Open(postgresDriver.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(&models.Album{})
	if err != nil {
		return
	}

	DB = database
}
