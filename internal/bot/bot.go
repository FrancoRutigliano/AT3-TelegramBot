package bot

import (
	"atomico3bot/config"
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func StartBot() error {
	token, err := config.SetUpConfig()
	if err != nil {
		log.Fatal("error loading .env data")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	return nil
}
