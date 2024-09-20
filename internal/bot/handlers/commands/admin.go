package commands

import (
	"atomico3bot/internal/auth"
	"errors"
	"fmt"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleBan(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	if !auth.IsAdmin(bot, message.Chat.ID, message.From.ID) {
		return errors.New("usuario no autorizado para ejecutar ese comando")
	}

	fmt.Println(message.ReplyToMessage)

	if message.ReplyToMessage == nil {
		return errors.New("debes responder a un mensaje del usuario que desees banear")
	}

	//id de usuario que deseemos banear
	userID := message.ReplyToMessage.From.ID

	chatMemberConfig := tgbotapi.ChatMemberConfig{
		ChatID: message.Chat.ID,
		UserID: userID,
	}

	_, err := bot.KickChatMember(tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: chatMemberConfig,
	})
	if err != nil {
		return errors.New("error al banear al usuario")
	}

	return nil
}
