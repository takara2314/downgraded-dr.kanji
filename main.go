package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/receive"
	"downgraded-dr.kanji/state"
	"downgraded-dr.kanji/utils"
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
	router.GET("/sheet.png", sheetGET)

	router.Run(":" + os.Getenv("PORT"))
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
			id := event.Source.UserID

			// Create state if not exist state.
			if _, exist := state.States[id]; !exist {
				state.States[id] = &state.State{}
			}

			// Detect rapid receive.
			diffMilli := int(time.Since(state.States[id].LastReceive) / time.Millisecond)
			if diffMilli <= 1000 {
				state.States[id].RapidCount += 1
			} else {
				state.States[id].RapidCount = 0
				state.States[id].IsRapidNotice = false
			}

			state.States[id].LastReceive = time.Now()

			// Get user profile to get user display name.
			profile, err := common.Bot.GetProfile(id).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}

			// Rapid notice.
			if state.States[id].RapidCount >= 5 {
				if !state.States[id].IsRapidNotice {
					fmt.Printf("[DETECT RAPID] %s\n", profile.DisplayName)
					state.States[id].IsRapidNotice = true

					_, err := common.Bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTextMessage(common.RapidNotice),
					).Do()
					if err != nil {
						log.Println(err)
						panic(err)
					}
				}

				return
			}

			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				fmt.Printf("[TEXT] %s: %s\n", profile.DisplayName, message.Text)
				receive.TextMessage(event, message.Text)

			case *linebot.ImageMessage:
				// Create a temporality file.
				filepath := filepath.Join(
					os.TempDir(),
					strconv.Itoa(utils.RandN(100000))+".png",
				)
				file, err := os.Create(filepath)
				if err != nil {
					log.Println(err)
					panic(err)
				}
				defer file.Close()

				// Get a image content.
				content, err := common.Bot.GetMessageContent(message.ID).Do()
				if err != nil {
					log.Println(err)
					panic(err)
				}
				defer content.Content.Close()

				// Write the bytes to the field.
				_, err = io.Copy(file, content.Content)
				if err != nil {
					log.Println(err)
					panic(err)
				}

				fmt.Printf("[IMAGE] %s\n", profile.DisplayName)
				receive.ImageMessage(event, filepath)
			}
		}
	}
}

func sheetGET(c *gin.Context) {
	c.File("./sheet.png")
}
