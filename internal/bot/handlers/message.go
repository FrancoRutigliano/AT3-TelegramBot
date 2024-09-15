package handlers

import (
	"atomico3bot/internal/filters"
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if filters.IsNegativeComment(message.Text) {

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
