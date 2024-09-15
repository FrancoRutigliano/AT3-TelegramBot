package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TokenTelegramAccess string
	TokenAccessHug      string
}

func SetUpConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		TokenTelegramAccess: os.Getenv("TOKEN_TELEGRAM_ACCESS"),
		TokenAccessHug:      os.Getenv("ACCESS_TOKEN_HUG"),
	}, nil
}
