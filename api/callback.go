package api

import (
	LineHandlers "go-line/Handlers"
	"go-line/Models"
	"go-line/Utils"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	channel_secret := os.Getenv("CHANNEL_SECRET")

	bot, err := linebot.New(channel_secret, channel_access_token)
	Utils.ErrorHandle(err)

	HandleLineEvent(r, bot)
}

func HandleLineEvent(r *http.Request, bot *linebot.Client) {

	events, _ := bot.ParseRequest(r)
	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {

			executor := LineHandlers.NewMessageHandler(
				&Models.HandleDto{
					Event:   event,
					Message: event.Message.(*linebot.TextMessage),
					Bot:     bot,
				})

			executor.Handle()
		}

	}
}
