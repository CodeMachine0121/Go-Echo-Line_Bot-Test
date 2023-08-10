package LineHandlers

import (
	"go-line/Models"
	serivces "go-line/Serivces"
	"go-line/Utils"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type SummerHandler struct {
	Dto *Models.HandleDto
}

func (s *SummerHandler) Handle() {

	transactionHistory := serivces.GetTransactionHistory()
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
