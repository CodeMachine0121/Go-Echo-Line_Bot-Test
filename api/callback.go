package api

import (
	LineHandlers "go-line/Handlers"
	"go-line/Models"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	channel_secret := os.Getenv("CHANNEL_SECRET")

	bot, _ := linebot.New(channel_secret, channel_access_token)

	executor := LineHandlers.NewMessageHandler(&Models.HandleDto{
		Bot: bot,
	})
	HandleLineEvent(w, r, executor)
}

func HandleLineEvent(w http.ResponseWriter, r *http.Request, executor *LineHandlers.MessageHandler) {
	events, _ := executor.Dto.Bot.ParseRequest(r)
	for _, event := range events {
		executor.Dto.Event = event

		if event.Type == linebot.EventTypeMessage {
			switch eventMessage := event.Message.(type) {
			case *linebot.TextMessage:
				executor.Dto.Message = eventMessage
				executor.Handle()
			}
		}
	}
}
