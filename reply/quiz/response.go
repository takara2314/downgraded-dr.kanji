package quiz

import (
	"downgraded-dr.kanji/common"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Response(event *linebot.Event, flexQuiz []byte) error {
	var err error
	var flexMessage linebot.FlexContainer

	flex, err := editFlex(flexQuiz)
	if err != nil {
		return err
	}

	flexMessage, err = linebot.UnmarshalFlexMessageJSON(flex)
	if err != nil {
		return err
	}

	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewFlexMessage(common.AnswerTheQuestion, flexMessage),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
