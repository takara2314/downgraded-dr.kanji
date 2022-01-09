package answer

import (
	"fmt"
	"strconv"
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

func synonym(event *linebot.Event, parameters []string) error {
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

	// Obtain content no.
	splitted := strings.Split(parameters[1], "_")
	// Require more than 1 parameter.
	if len(splitted) < 1 {
		err := invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	baseOfContentNo := strings.Split(splitted[0], ",")
	// Require 2 parameters.
	if len(baseOfContentNo) != 2 {
		err := invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	contentNo, err := utils.AtoiSlice(baseOfContentNo)
	if err != nil {
		err := invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	// Make a content.
	content := fmt.Sprintf(
		"%s ≒ %s",
		common.Quizzes.Synonyms[no-1][contentNo[0]-1],
		common.Quizzes.Synonyms[no-1][contentNo[1]-1],
	)

	// If its quiz no has more than 3 contents, answer the all.
	if len(common.Quizzes.Synonyms[no-1]) > 2 {
		// Make more content.
		moreContent := strings.Join(common.Quizzes.Synonyms[no-1], "、")

		// Reply the content.
		_, err = common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(
				fmt.Sprintf("%s\n%s", common.SuggestedAnswerTitle, content),
			),
			linebot.NewTextMessage(
				fmt.Sprintf("%s\n%s", common.MoreSynonyms, moreContent),
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
