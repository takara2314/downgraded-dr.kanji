package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	var err error

	bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	router := gin.Default()
	router.POST("/callback", callbackPOST)

	var port string = os.Getenv("PORT")
	router.Run(":" + port)
}

func callbackPOST(c *gin.Context) {
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				profile, err := bot.GetProfile(event.Source.UserID).Do()
				if err != nil {
					log.Println(err)
					panic(err)
				}
				fmt.Printf("[CHAT] %s: %s\n", profile.DisplayName, message.Text)
				replyTextMessage(event, message.Text)
			}
		}
	}
}
