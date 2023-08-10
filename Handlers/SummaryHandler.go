package LineHandlers

import (
	"go-line/Models"
	singletons "go-line/Singletons"
	"go-line/Utils"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type SummerHandler struct {
	Dto *Models.HandleDto
}

func (s *SummerHandler) Handle() {

	// TODO: use singleton get history
	transactionHistory := singletons.GetTransactionSingleton()
	for _, t := range transactionHistory.History {

		replyMsg := Utils.GetTimeWithFormat(t.CreatedTime) + " " + t.Item + " " + strconv.Itoa(t.Amount)

		_, err := s.Dto.Bot.ReplyMessage(s.Dto.Event.ReplyToken, linebot.NewTextMessage(replyMsg)).Do()

		Utils.ErrorHandle(err)
	}
	s.Dto.Bot.ReplyMessage(s.Dto.Event.ReplyToken, linebot.NewTextMessage("Total: "+strconv.Itoa(transactionHistory.Totals)))
}

func NewSummerHandler(dto *Models.HandleDto) *SummerHandler {
	return &SummerHandler{Dto: dto}
}
