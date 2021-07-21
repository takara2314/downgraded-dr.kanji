package quiz

import "github.com/line/line-bot-sdk-go/linebot"

func Response(bot *linebot.Client, event *linebot.Event, flexQuiz []byte) error {
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

	_, err = bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewFlexMessage("問題に答えよ。", flexMessage),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
