package handlers

import (
	"atomico3bot/internal/service"
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, accessToken string) {

	isNegative, err := service.IsNegative(message.Text, accessToken)
	if err != nil {
		log.Fatal("error: ", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "error analizando el mensaje")
		bot.Send(msg)
		return
	}
	if isNegative {
		deleteMessage := tgbotapi.DeleteMessageConfig{
			ChatID:    message.Chat.ID,
			MessageID: message.MessageID,
		}

		_, err := bot.DeleteMessage(deleteMessage)
		if err != nil {
			log.Fatal("error al eliminar un comentario negativo")
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "Por favor evita malos comentarios")
		bot.Send(msg)

	}

}
