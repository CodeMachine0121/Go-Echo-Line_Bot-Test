package LineHandlers

import (
	"go-line/Models"
	serivces "go-line/Serivces"
	"go-line/Utils"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type AccountingHandler struct {
	Dto *Models.HandleDto
}

func (a *AccountingHandler) Handle() {
	// TODO: should contain 3 elements: "count", item, count
	splitCommand := strings.Split(a.Dto.Message.Text, " ")
	amount, _ := strconv.Atoi(splitCommand[2])
	totalSpend := serivces.InsertTransaction(&Models.SigleTransaction{
		Id:          uuid.New(),
		Item:        splitCommand[1],
		Amount:      amount,
		CreatedTime: time.Now(),
	})

	_, err := a.Dto.Bot.ReplyMessage(a.Dto.Event.ReplyToken, linebot.NewTextMessage(strconv.Itoa(totalSpend))).Do()

	Utils.ErrorHandle(err)
}

func NewAccountingHandler(dto *Models.HandleDto) *AccountingHandler {
	return &AccountingHandler{Dto: dto}
}
