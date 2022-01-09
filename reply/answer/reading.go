package answer

import (
	"fmt"
	"strconv"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

func reading(event *linebot.Event, parameters []string) error {
	// Require 3 parameter.
	if len(parameters) != 3 {
		err := invalid(event)
		if err != nil {
			return err
		}
		return nil
	}

	switch parameters[0] {
	case "antonyms":
		err := readingFromAntonyms(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "homonyms":
		err := readingFromHomonyms(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "synonyms":
		err := readingFromSynonyms(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "confers":
		err := readingFromConfers(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "others":
		err := readingFromOthers(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	default:
		err := invalid(event)
		if err != nil {
			return err
		}
	}

	return nil
}

func readingFromAntonyms(event *linebot.Event, parameters []string) error {
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

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s",
				common.SuggestedAnswerTitle,
				obtainYomiMoji(common.Quizzes.Antonyms[no-1][contentNo-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func readingFromHomonyms(event *linebot.Event, parameters []string) error {
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

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s",
				common.SuggestedAnswerTitle,
				common.Quizzes.Homonyms[no-1][0],
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func readingFromSynonyms(event *linebot.Event, parameters []string) error {
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

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s",
				common.SuggestedAnswerTitle,
				obtainYomiMoji(common.Quizzes.Synonyms[no-1][contentNo-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func readingFromConfers(event *linebot.Event, parameters []string) error {
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

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s",
				common.SuggestedAnswerTitle,
				obtainYomiMoji(common.Quizzes.Confers[no-1][contentNo-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func readingFromOthers(event *linebot.Event, parameters []string) error {
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

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s",
				common.SuggestedAnswerTitle,
				obtainYomiMoji(common.Quizzes.Others[no-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

// obtainYomiMoji returns the yomi only.
func obtainYomiMoji(s string) string {
	return utils.ObtainStringCovered(s, "（", "）")
}
