package handler

import (
	Handlers "go-line/Hamdlers"
	Models "go-line/Models"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client
var msgHandler *Handlers.MessageHandler

func init() {
	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	channel_secret := os.Getenv("CHANNEL_SECRET")
	bot, _ = linebot.New(channel_secret, channel_access_token)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	events, _ := bot.ParseRequest(r)

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch eventMessage := event.Message.(type) {

			case *linebot.TextMessage:
				msgHandler.Handle(
					&Models.HandleDto{
						Event:   *event,
						Message: *eventMessage,
						Bot:     *bot,
					})
			}
		}
	}

}
