package receive

import (
	"fmt"
	"log"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/state"
	"downgraded-dr.kanji/utils"

	"github.com/line/line-bot-sdk-go/linebot"
)

func ImageMessage(event *linebot.Event, filepath string) {
	// May receive a quiz answer.
	if state.States[event.Source.UserID].IsQuizzing {
		state.States[event.Source.UserID].IsQuizzing = false

		// Load a image file.
		file, _ := os.Open(filepath)
		defer file.Close()
		image, _ := vision.NewImageFromReader(file)

		// Detect texts from image.
		texts, err := common.VisionAPI.DetectTexts(
			common.VisionAPICtx,
			image,
			nil,
			10,
		)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		detectedText := ""

		// Detected something.
		if len(texts) > 0 {
			detectedText = texts[0].Description

			// Trim unrelated chars.
			detectedText = strings.Replace(detectedText, " ", "", -1)
			detectedText = strings.Replace(detectedText, "　", "", -1)
			detectedText = strings.Replace(detectedText, "\n", "", -1)

			fmt.Println("[TEXT DETECTED]", detectedText)
		}

		// Detected nothing.
		if detectedText == "" {
			_, err := common.Bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage(common.DetectedNothing),
			).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}
			return
		}

		// It is the correct answer if contains $~.LastQuiz.Corrects.
		if utils.StringSliceFind(
			state.States[event.Source.UserID].LastQuiz.Corrects,
			detectedText,
		) != -1 {
			// Correct answered.
			_, err := common.Bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage(
					fmt.Sprintf("%s: %s", common.DetectedText, detectedText),
				),
				linebot.NewTextMessage(common.YourAnswerCorrect),
			).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}

		} else {
			// Wrong answered.
			_, err := common.Bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage(
					fmt.Sprintf("%s: %s", common.DetectedText, detectedText),
				),
				linebot.NewTextMessage(common.YourAnswerWrong),
			).Do()
			if err != nil {
				log.Println(err)
				panic(err)
			}
		}

		return
	}

	_, err := common.Bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage(common.UnknownImage),
	).Do()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
