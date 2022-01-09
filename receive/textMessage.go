package receive

import (
	"log"
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/reply/answer"
	"downgraded-dr.kanji/reply/quiz"
	"downgraded-dr.kanji/state"
	"downgraded-dr.kanji/utils"

	"github.com/line/line-bot-sdk-go/linebot"
)

func TextMessage(event *linebot.Event, message string) {
	if strings.HasPrefix(message, "quiz") {
		// Receive a text "quiz **"
		err := quiz.Response(
			event,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "answer") {
		// Receive a text "answer **"
		state.States[event.Source.UserID].IsQuizzing = false

		err := answer.Response(
			event,
			message,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if message == "sheet" {
		// Receive a text "sheet"
		_, err := common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewImageMessage(
				common.ServiceURL+"/sheet.png",
				common.ServiceURL+"/sheet.png",
			),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else {
		// Receive a else text of above.

		// May receive a quiz answer.
		if state.States[event.Source.UserID].IsQuizzing {
			state.States[event.Source.UserID].IsQuizzing = false

			// It is the correct answer if contains $~.LastQuiz.Corrects.
			if utils.StringSliceFind(
				state.States[event.Source.UserID].LastQuiz.Corrects,
				message,
			) != -1 {
				// Correct answered.
				_, err := common.Bot.ReplyMessage(
					event.ReplyToken,
					linebot.NewTextMessage(common.YourAnswerCorrect),
				).Do()
				if err != nil {
					log.Println(err)
					panic(err)
				}

			} else {
				// Wrong answered.
				_, err := common.Bot.ReplyMessage(
					event.ReplyToken,
					linebot.NewTextMessage(common.YourAnswerWrong),
				).Do()
				if err != nil {
					log.Println(err)
					panic(err)
				}
			}

			return
		}

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
