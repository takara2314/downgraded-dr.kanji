package quiz

import (
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/state"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Response(event *linebot.Event, message string) error {
	splitted := strings.Split(message, " ")
	if len(splitted) <= 1 {
		splitted = make([]string, 2)
	}

	// Choice a quiz.
	quiz, err := choice(splitted)
	if err != nil {
		_, err = common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(common.InvalidParameter),
		).Do()
		if err != nil {
			return err
		}
	}

	// Set a state
	state.States[event.Source.UserID].LastQuiz = &quiz
	state.States[event.Source.UserID].IsQuizzing = true

	// Create JSON to send a flex message.
	flex, err := createFlexMessage(quiz)
	if err != nil {
		return err
	}

	// Convert to flex container from JSON.
	container, err := linebot.UnmarshalFlexMessageJSON(flex)
	if err != nil {
		return err
	}

	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewFlexMessage(common.AnswerTheQuestion, container),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
