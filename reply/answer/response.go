package answer

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Response(event *linebot.Event, message string) error {
	splitted := strings.Split(message, " ")

	switch splitted[1] {
	case "Antonym":
		err := antonym(event, splitted[2:])
		if err != nil {
			return err
		}

	case "Homonym":
		err := homonym(event, splitted[2:])
		if err != nil {
			return err
		}

	case "Synonym":
		err := synonym(event, splitted[2:])
		if err != nil {
			return err
		}

	// case "Confer":
	// 	err := confer(event, splitted[2:])
	// 	if err != nil {
	// 		return err
	// 	}

	// case "Writing":
	// 	err := writing(event, splitted[2:])
	// 	if err != nil {
	// 		return err
	// 	}

	// case "Reading":
	// 	err := reading(event, splitted[2:])
	// 	if err != nil {
	// 		return err
	// 	}

	default:
		err := invalid(event)
		if err != nil {
			return err
		}
	}

	// var err error

	// splited := strings.Split(message, " ")

	// if len(splited) != 3 {
	// 	_, err = common.Bot.ReplyMessage(
	// 		event.ReplyToken,
	// 		linebot.NewTextMessage(common.InvalidCommand),
	// 	).Do()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// // 問題集
	// var quizes [][]string
	// // つなぎ記号
	// var joinChar string = "，"

	// switch splited[1] {
	// case "Antonyms":
	// 	quizes = common.Quizzes.Antonyms
	// 	joinChar = "←→ "
	// case "Homonym":
	// 	quizes = common.Quizzes.Homonyms
	// case "Synonyms":
	// 	quizes = common.Quizzes.Synonyms
	// 	joinChar = "≒ "
	// case "Confer":
	// 	quizes = common.Quizzes.Confers
	// default:
	// 	_, err = common.Bot.ReplyMessage(
	// 		event.ReplyToken,
	// 		linebot.NewTextMessage(common.InvalidCommand),
	// 	).Do()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// index, err := strconv.Atoi(splited[2])
	// if err != nil {
	// 	_, err = common.Bot.ReplyMessage(
	// 		event.ReplyToken,
	// 		linebot.NewTextMessage(common.InvalidCommand),
	// 	).Do()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// if index >= len(quizes) {
	// 	_, err = common.Bot.ReplyMessage(
	// 		event.ReplyToken,
	// 		linebot.NewTextMessage(common.InvalidCommand),
	// 	).Do()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// // _, err = bot.ReplyMessage(
	// // 	event.ReplyToken,
	// // 	linebot.NewTextMessage(
	// // 		strings.Join(
	// // 			quizes[index][1:len(quizes[index])],
	// // 			"，",
	// // 		),
	// // 	),
	// // ).Do()
	// // if err != nil {
	// // 	return err
	// // }

	// _, err = common.Bot.ReplyMessage(
	// 	event.ReplyToken,
	// 	linebot.NewTextMessage(
	// 		"【問題と答え】\n"+
	// 			strings.Join(
	// 				quizes[index],
	// 				joinChar,
	// 			),
	// 	),
	// ).Do()
	// if err != nil {
	// 	return err
	// }

	return nil
}
