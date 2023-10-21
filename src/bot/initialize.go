package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gobuffalo/envy"
)

func Initialize() (*tgbotapi.BotAPI, error) {
	token, err := envy.MustGet("TG_BOT_TOKEN")
	if err != nil {
		panic(err)
	}
	bot, errBot := tgbotapi.NewBotAPI(token)
	if errBot != nil {
		return nil, errBot
	}
	return bot, nil
}
