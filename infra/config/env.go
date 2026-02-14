package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host string
	Port string
}

func LoadEnv() *Env {
	_ = godotenv.Load()

	return &Env{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
