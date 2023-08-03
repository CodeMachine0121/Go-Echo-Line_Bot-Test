package handlers

import (
	Models "go-line/Models"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MessageHandler struct{}

func (m *MessageHandler) Handle(dto *Models.HandleDto) {
	_, err := dto.Bot.ReplyMessage(dto.Event.ReplyToken, linebot.NewTextMessage(dto.Message.Text)).Do()
	if err != nil {
		log.Fatal(err)
	}
}
