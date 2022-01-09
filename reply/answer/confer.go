package answer

import (
	"fmt"
	"strconv"
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

func confer(event *linebot.Event, parameters []string) error {
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
	splitted := strings.Split(parameters[1], "_")
	// Require more than 1 parameter.
	if len(splitted) < 1 {
		err := invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	// If $contentNo is not completely string, the bot stops process.
	contentNo, err := strconv.Atoi(splitted[0])
	if err != nil {
		err := invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	// Make a content.
	content := common.Quizzes.Confers[no-1][contentNo]

	// If its quiz no has more than 3 contents, answer the all.
	if len(common.Quizzes.Synonyms[no-1]) > 2 {
		// Make more content.
		others := utils.StringSliceRemove(common.Quizzes.Confers[no-1], contentNo)
		others = utils.StringSliceRemove(others, 0)
		moreContent := strings.Join(others, "、")

		// Reply the content.
		_, err = common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(
				fmt.Sprintf("%s\n%s", common.SuggestedAnswerTitle, content),
			),
			linebot.NewTextMessage(
				fmt.Sprintf("%s\n%s", common.MoreConfers, moreContent),
			),
		).Do()
		if err != nil {
			return err
		}

	} else {
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
	}

	return nil
}
