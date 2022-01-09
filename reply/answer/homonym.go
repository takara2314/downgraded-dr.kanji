package answer

import (
	"fmt"
	"strconv"
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

func homonym(event *linebot.Event, parameters []string) error {
	// Require 2 parameters.
	if len(parameters) != 2 {
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

	// Obtain a content no.
	// If $no is not completely string, the bot stops process.
	contentNo, err := strconv.Atoi(parameters[1])
	if err != nil {
		err = invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	// Make more content.
	quizzesExceptContentNo := utils.StringSliceRemove(common.Quizzes.Homonyms[no-1], contentNo)
	quizzesExceptContentNo = utils.StringSliceRemove(quizzesExceptContentNo, 0)
	content := strings.Join(quizzesExceptContentNo, "、")

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s %s\n%s: %s",
				common.SuggestedAnswerTitle,
				content,
				common.HomonymOneOfThese,
				common.HomonymYomi,
				common.Quizzes.Homonyms[no-1][0],
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
