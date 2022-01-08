package receive

import (
	"log"

	"downgraded-dr.kanji/common"

	"github.com/line/line-bot-sdk-go/linebot"
)

func ImageMessage(event *linebot.Event, message string) {
	_, err := common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage("？"),
	).Do()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
