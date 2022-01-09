package answer

import (
	"fmt"
	"strconv"
	"strings"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
	"github.com/line/line-bot-sdk-go/linebot"
)

func writing(event *linebot.Event, parameters []string) error {
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
		err := writingFromAntonyms(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "homonyms":
		err := writingFromHomonyms(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "synonyms":
		err := writingFromSynonyms(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "confers":
		err := writingFromConfers(event, parameters[1:])
		if err != nil {
			err = invalid(event)
			if err != nil {
				return err
			}
		}

	case "others":
		err := writingFromOthers(event, parameters[1:])
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

func writingFromAntonyms(event *linebot.Event, parameters []string) error {
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
				snipYomiMoji(common.Quizzes.Antonyms[no-1][contentNo-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func writingFromHomonyms(event *linebot.Event, parameters []string) error {
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

	contents := strings.Join(common.Quizzes.Homonyms[no-1][1:], "、")

	// Reply the content.
	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			fmt.Sprintf("%s\n%s %s",
				common.SuggestedAnswerTitle,
				contents,
				common.HomonymOneOfThese,
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func writingFromSynonyms(event *linebot.Event, parameters []string) error {
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
				snipYomiMoji(common.Quizzes.Synonyms[no-1][contentNo-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func writingFromConfers(event *linebot.Event, parameters []string) error {
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
				snipYomiMoji(common.Quizzes.Confers[no-1][contentNo-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

func writingFromOthers(event *linebot.Event, parameters []string) error {
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
				snipYomiMoji(common.Quizzes.Others[no-1]),
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}

// snipYomiMoji returns the moji snipped yomi.
func snipYomiMoji(s string) string {
	return utils.SnipStringCovered(s, "（", "）")
}
