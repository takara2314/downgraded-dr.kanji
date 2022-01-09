package receive

import (
	"log"
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/reply/quiz"

	"github.com/line/line-bot-sdk-go/linebot"
)

func TextMessage(event *linebot.Event, message string) {
	if strings.HasPrefix(message, "quiz") {
		// Receive a text "quiz **"
		err := quiz.Response(
			event,
			common.FlexQuiz,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		// } else if strings.HasPrefix(message, "answer") {
		// 	// Receive a text "answer **"
		// 	err := answer.Response(
		// 		event,
		// 		message,
		// 	)
		// 	if err != nil {
		// 		log.Println(err)
		// 		panic(err)
		// 	}

	} else {
		// Receive a else text of above
		_, err := common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(common.UnknownMessage),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
