package main

import (
	"go-line/Utils"
	"go-line/api"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	channel_secret := os.Getenv("CHANNEL_SECRET")

	bot, err := linebot.New(channel_secret, channel_access_token)
	Utils.ErrorHandle(err)

	http.HandleFunc("/api/callback", func(w http.ResponseWriter, r *http.Request) {
		api.HandleLineEvent(r, bot)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
