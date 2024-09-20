package handlers

import (
	"atomico3bot/internal/bot/handlers/commands"
	"atomico3bot/internal/service"
	"fmt"
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

	case "rules":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Reglas del grupo: \n1. Respeto mutuo \n2. No spam \n3. No mensajes ofensivos")
		bot.Send(msg)

	case "ban":
		if err := commands.HandleBan(bot, message); err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
			bot.Send(msg)
			return
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "usuario baneado correctamente")
		bot.Send(msg)

	case "price":

		price, err := service.GetTokenPrice()
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Lo siento, no pude obtener el precio del token en este momento.")
			bot.Send(msg)
			return
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("El precio actual de Atomico es: %.6f USDT", price))
		bot.Send(msg)

	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Comando no reconocido.")
		bot.Send(msg)
	}
}
