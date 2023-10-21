package main

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/madiyarrakhman/nice-meet/internal/infrastructure"
	"github.com/madiyarrakhman/nice-meet/internal/models"
	"log"
)

var commands = map[string]string{
	"start":  "Start bot",
	"help":   "Help",
	"status": "Get status",
	"ping":   "Ping",
}

func main() {
	app := infrastructure.NewApp()

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := app.Bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "start":
				errAddMember := app.MemberService.AddMember(
					update.Message.From.ID,
					update.Message.From.UserName,
					update.Message.From.FirstName,
					update.Message.From.LastName,
				)

				if errors.Is(errAddMember, models.ErrMemberExistAlready) {
					msg.Text = "Member exist already"
					break
				}

				if errAddMember != nil {
					msg.Text = "Error adding member"
					log.Default().Println(errAddMember)
					break
				}
			}
			_, errSend := app.Bot.Send(msg)
			if errSend != nil {
				log.Default().Println(errSend)
			}
		}
	}
}
