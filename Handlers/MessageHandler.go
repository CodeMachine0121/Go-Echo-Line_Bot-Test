package LineHandlers

import (
	"go-line/Models"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MessageHandler struct {
	Dto *Models.HandleDto
}

func (m *MessageHandler) Handle() {
	_, err := m.Dto.Bot.ReplyMessage(m.Dto.Event.ReplyToken, linebot.NewTextMessage(m.Dto.Message.Text)).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func NewMessageHandler(dto *Models.HandleDto) *MessageHandler {
	return &MessageHandler{
		Dto: dto,
	}
}
