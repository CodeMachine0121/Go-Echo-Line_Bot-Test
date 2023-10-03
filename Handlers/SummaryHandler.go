package LineHandlers

import (
	"encoding/json"
	"go-line/Models"
	singletons "go-line/Singletons"
	"go-line/Utils"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type SummerHandler struct {
	Dto *Models.HandleDto
}

// TODO: use singleton get history

func (s *SummerHandler) Handle() {
	HandleOnGettingHistory(s)
	HandleOnClearingHistory(s)
}

func HandleOnClearingHistory(s *SummerHandler) {
	transactionHistory := singletons.GetTransactionSingleton()
	transactionHistory.InitProperty()
}

func NewSummerHandler(dto *Models.HandleDto) *SummerHandler {
	return &SummerHandler{Dto: dto}
}

func HandleOnGettingHistory(s *SummerHandler) {
	transactionHistory := singletons.GetTransactionSingleton()

	transactionHistorySerialized, err := json.Marshal(transactionHistory)
	Utils.ErrorHandle(err)

	if _, err := s.Dto.Bot.ReplyMessage(s.Dto.Event.ReplyToken, linebot.NewTextMessage(string(transactionHistorySerialized))).Do(); err != nil {
		Utils.ErrorHandle(err)
	}

	Utils.ErrorHandle(err)
	s.Dto.Bot.ReplyMessage(s.Dto.Event.ReplyToken, linebot.NewTextMessage("Total: "+strconv.Itoa(transactionHistory.Totals))).Do()
}
