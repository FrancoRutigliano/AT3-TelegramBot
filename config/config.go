package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TokenTelegramAccess string
}

func SetUpConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		TokenTelegramAccess: os.Getenv("TOKEN_TELEGRAM_ACCESS"),
	}, nil
}
