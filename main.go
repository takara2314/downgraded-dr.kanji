package main

import (
	"fmt"
	"log"
	"os"

	"downgraded-dr.kanji/common"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	var err error

	common.Bot, err = linebot.New(
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
	events, err := common.Bot.ParseRequest(c.Request)
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
			// Get user profile to get user display name
			profile, err := common.Bot.GetProfile(event.Source.UserID).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}

			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				fmt.Printf("[TEXT] %s: %s\n", profile.DisplayName, message.Text)
				receive.TextMessage(event, message.Text)

			case *linebot.ImageMessage:
				fmt.Printf("[IMAGE] %s %s\n", profile.DisplayName, message.OriginalContentURL)
				receive.ImageMessage(event, message.OriginalContentURL)
			}
		}
	}
}
