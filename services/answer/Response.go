package answer

import (
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Response(bot *linebot.Client, event *linebot.Event, message string) error {
	var err error

	splited := strings.Split(message, " ")

	if len(splited) != 3 {
		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage("不正なコマンドです。"),
		).Do()
		if err != nil {
			return err
		}
	}

	// 問題集
	var quizes [][]string

	switch splited[1] {
	case "Antonyms":
		quizes = Config.Antonyms
	case "Homonym":
		quizes = Config.Homonym
	case "Synonyms":
		quizes = Config.Synonyms
	case "Confer":
		quizes = Config.Confer
	case "Four":
		quizes = Config.Four
	default:
		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage("不正なコマンドです。"),
		).Do()
		if err != nil {
			return err
		}
	}

	index, err := strconv.Atoi(splited[2])
	if err != nil {
		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage("不正なコマンドです。"),
		).Do()
		if err != nil {
			return err
		}
	}

	if index >= len(quizes) {
		_, err = bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage("不正なコマンドです。"),
		).Do()
		if err != nil {
			return err
		}
	}

	_, err = bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(
			strings.Join(
				quizes[index][1:len(quizes[index])],
				"，",
			),
		),
	).Do()
	if err != nil {
		return err
	}

	return nil
}
