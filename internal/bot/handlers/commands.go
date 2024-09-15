package handlers

import (
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleCommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		log.Println("soy el comando start")
		msg := tgbotapi.NewMessage(message.Chat.ID, "¡Hola! Soy un bot de moderación de Atomico 3, en que puedo ayudarte?")
		bot.Send(msg)

	case "help":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Comandos disponibles:\n/start - Iniciar el bot\n/help - Mostrar ayuda\n/info - información sobre atomico\n/stacking - información sobre como hacer stacking\n/price - precio actual de la Atomico 3")
		bot.Send(msg)

	case "info":

	case "stacking":

	case "price":

	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Comando no reconocido.")
		bot.Send(msg)
	}
}
