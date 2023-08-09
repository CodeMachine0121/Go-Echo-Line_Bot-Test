package LineHandlers

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go-line/Models"
	"go-line/Utils"
)

type MessageHandler struct {
	Dto *Models.HandleDto
}

func (m *MessageHandler) Handle() {
	_, err := m.Dto.Bot.ReplyMessage(m.Dto.Event.ReplyToken, linebot.NewTextMessage(m.Dto.Message.Text)).Do()
	Utils.ErrorHandle(err)
}

func NewMessageHandler(dto *Models.HandleDto) *MessageHandler {
	return &MessageHandler{
		Dto: dto,
	}
}
