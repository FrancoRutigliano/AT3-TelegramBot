package config

import (
	"os"

	"github.com/joho/godotenv"
)

func SetUpConfig() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	return os.Getenv("TOKEN_ACCESS"), nil
}
