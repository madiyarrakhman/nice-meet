package src

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gobuffalo/envy"
	bot2 "nice-meet/src/bot"
	"nice-meet/src/repository"
	"nice-meet/src/services"
)

type App struct {
	Bot           *tgbotapi.BotAPI
	MemberService *services.Member
}

func NewApp() *App {
	err := envy.Load()
	if err != nil {
		panic(err)
	}

	bot, errBot := bot2.Initialize()
	if errBot != nil {
		panic(errBot)
	}

	db, err := repository.Initialize()
	if err != nil {
		panic(err)
	}

	memberService := services.NewMember(repository.NewMember(db))
	return &App{
		MemberService: memberService,
		Bot:           bot,
	}
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop() error {
	return nil
}
