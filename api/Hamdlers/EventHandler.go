package handler

import "github.com/line/line-bot-sdk-go/v7/linebot"

type LineEventHandler interface {
	Handle()
}

type HandleDto struct {
	Event   *linebot.Event
	Message *linebot.TextMessage
	Bot     *linebot.Client
}
