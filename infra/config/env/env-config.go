package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host       string
	Port       string
	JwtSecret  string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func LoadEnv() *Env {
	_ = godotenv.Load()

	return &Env{
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
