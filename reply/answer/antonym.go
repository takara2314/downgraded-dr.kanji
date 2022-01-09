package answer

import (
	"fmt"
	"strconv"

	"downgraded-dr.kanji/common"
	"github.com/line/line-bot-sdk-go/linebot"
)

func antonym(event *linebot.Event, parameters []string) error {
	// Require more than 1 parameter.
	if len(parameters) < 1 {
		err := invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	// Obtain a quiz no.
	// If $no is not completely string, the bot stops process.
	no, err := strconv.Atoi(parameters[0])
	if err != nil {
		err = invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	// Make a content.
	content := fmt.Sprintf(
		"%s ←→ %s",
		common.Quizzes.Antonyms[no-1][0],
		common.Quizzes.Antonyms[no-1][1],
	)

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s", common.SuggestedAnswerTitle, content),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
