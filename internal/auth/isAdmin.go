package auth

import tgbotapi "gopkg.in/telegram-bot-api.v4"

func IsAdmin(bot *tgbotapi.BotAPI, chatID int64, userID int) bool {
	role, err := bot.GetChatMember(tgbotapi.ChatConfigWithUser{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
	}

	if role.Status == "administrator" || role.Status == "creator" {
		return true
	}

	return false

}
