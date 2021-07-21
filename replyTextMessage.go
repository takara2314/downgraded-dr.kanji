package main

import (
	"log"
	"strings"
	"tcj3-kadai-tuika-kun/services/answer"
	"tcj3-kadai-tuika-kun/services/quiz"

	"github.com/line/line-bot-sdk-go/linebot"
)

func replyTextMessage(event *linebot.Event, message string) {
	if strings.HasPrefix(message, "quiz") {
		err := quiz.Response(
			bot,
			event,
			flexQuiz,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

	} else if strings.HasPrefix(message, "answer") {
		err := answer.Response(
			bot,
			event,
			message,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
