package receive

import (
	"log"
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/reply/answer"
	"downgraded-dr.kanji/reply/quiz"

	"github.com/line/line-bot-sdk-go/linebot"
)

func TextMessage(event *linebot.Event, message string) {
	if strings.HasPrefix(message, "quiz") {
		// Response "quiz **"
		err := quiz.Response(
			event,
			common.FlexQuiz,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "answer") {
		// Response "answer **"
		err := answer.Response(
			event,
			message,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else {
		// Response else of above
		_, err := common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage("現在quizと一部のコマンド以外対応できません。\nたからーんの貴重なテスト勉強時間にワシが呼び戻されたもの。"),
		).Do()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
