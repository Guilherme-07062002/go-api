package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host      string
	Port      string
	JwtSecret string
}

func LoadEnv() *Env {
	_ = godotenv.Load()

	return &Env{
		Host:      os.Getenv("HOST"),
		Port:      os.Getenv("PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
