package bot

import (
	"atomico3bot/config"
	"atomico3bot/internal/bot/handlers"
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func StartBot() error {
	config, err := config.SetUpConfig()
	if err != nil {
		log.Fatal("error loading .env data")
	}

	bot, err := tgbotapi.NewBotAPI(config.TokenTelegramAccess)
	if err != nil {
		return err
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal("error al obtener actualizaciones")
	}

	for update := range updates {
		if update.Message != nil {
			log.Printf("Mensaje recibido: %s", update.Message.Text)
			log.Printf("Mensaje de respuesta: %v", update.Message.ReplyToMessage)

			if update.Message.IsCommand() {
				handlers.HandleCommands(bot, update.Message)
				continue
			}

			handlers.HandleMessage(bot, update.Message, config.TokenAccessHug)
		}
	}

	return nil
}
