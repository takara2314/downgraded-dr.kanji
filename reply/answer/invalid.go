package answer

import (
	"downgraded-dr.kanji/common"
	"github.com/line/line-bot-sdk-go/linebot"
)

func invalid(event *linebot.Event) error {
	_, err := common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(common.InvalidCommand),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
