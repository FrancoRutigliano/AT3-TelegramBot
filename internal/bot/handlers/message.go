package handlers

import (
	"atomico3bot/internal/filters"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if filters.IsNegativeComment(message.Text) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Por favor evita malos comentarios")
		bot.Send(msg)
	}
}
