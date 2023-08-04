package Models

import "github.com/line/line-bot-sdk-go/v7/linebot"

type HandleDto struct {
	Event   *linebot.Event
	Message *linebot.TextMessage
	Bot     *linebot.Client
}
