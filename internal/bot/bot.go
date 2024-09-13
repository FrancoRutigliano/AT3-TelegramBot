package bot

import (
	"atomico3bot/config"
	"log"
)

func StartBot() {
	_, err := config.SetUpConfig()
	if err != nil {
		log.Fatal("error loading .env data")
	}
}
