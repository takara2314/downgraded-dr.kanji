package answer

import (
	"strconv"
	"strings"

	"downgraded-dr.kanji/common"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Response(event *linebot.Event, message string) error {
	var err error

	splited := strings.Split(message, " ")

	if len(splited) != 3 {
		_, err = common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(common.InvalidCommand),
		).Do()
		if err != nil {
			return err
		}
	}

	// 問題集
	var quizes [][]string
	// つなぎ記号
	var joinChar string = "，"

	switch splited[1] {
	case "Antonyms":
		quizes = common.Quizzes.Antonyms
		joinChar = "←→ "
	case "Homonym":
		quizes = common.Quizzes.Homonyms
	case "Synonyms":
		quizes = common.Quizzes.Synonyms
		joinChar = "≒ "
	case "Confer":
		quizes = common.Quizzes.Confers
	default:
		_, err = common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(common.InvalidCommand),
		).Do()
		if err != nil {
			return err
		}
	}

	index, err := strconv.Atoi(splited[2])
	if err != nil {
		_, err = common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(common.InvalidCommand),
		).Do()
		if err != nil {
			return err
		}
	}

	if index >= len(quizes) {
		_, err = common.Bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(common.InvalidCommand),
		).Do()
		if err != nil {
			return err
		}
	}

	// _, err = bot.ReplyMessage(
	// 	event.ReplyToken,
	// 	linebot.NewTextMessage(
	// 		strings.Join(
	// 			quizes[index][1:len(quizes[index])],
	// 			"，",
	// 		),
	// 	),
	// ).Do()
	// if err != nil {
	// 	return err
	// }

	_, err = common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			"【問題と答え】\n"+
				strings.Join(
					quizes[index],
					joinChar,
				),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
