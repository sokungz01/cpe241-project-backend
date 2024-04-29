package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	FRONTEND_URL string
	DB_URL       string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	frontURL := os.Getenv("FRONTEND_URL")
	dbURL := os.Getenv("DB_URL")

	return &Config{
		FRONTEND_URL: frontURL,
		DB_URL:       dbURL,
	}, nil
}
