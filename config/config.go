package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	FRONTEND_URL     string
	DB_URL           string
	JWT_ACCESS_TOKEN string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	frontURL := os.Getenv("FRONTEND_URL")
	dbURL := os.Getenv("DB_URL")
	jwtToken := os.Getenv("JWT_ACCESS_TOKEN")

	return &Config{
		FRONTEND_URL:     frontURL,
		DB_URL:           dbURL,
		JWT_ACCESS_TOKEN: jwtToken,
	}, nil
}
