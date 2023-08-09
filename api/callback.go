package api

import (
	LineHandlers "go-line/Handlers"
	"go-line/Models"
	serivces "go-line/Serivces"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	channel_secret := os.Getenv("CHANNEL_SECRET")

	bot, err := linebot.New(channel_secret, channel_access_token)

	if err != nil {
		log.Fatal("linebot create fail", err)
	}

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
				if strings.Contains(eventMessage.Text, "echo") {
					executor.Dto.Message = eventMessage
					executor.Handle()
				}
				if strings.Contains(eventMessage.Text, "count") {
					// TODO: should contain 3 elements: "count", item, count
					command := strings.Split(eventMessage.Text, " ")
					spend, _ := strconv.Atoi(command[1])
					totalSpend := serivces.InserTransaction(&Models.SigleTransaction{
						Id:     uuid.New(),
						Item:   command[2],
						Amount: spend,
					})
					_, err := executor.Dto.Bot.ReplyMessage(executor.Dto.Event.ReplyToken, linebot.NewTextMessage(strconv.Itoa(totalSpend))).Do()

					if err != nil {
						log.Fatal(err)
					}

				}
			}

		}
	}
}
