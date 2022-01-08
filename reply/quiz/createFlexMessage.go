package quiz

import (
	"errors"

	"downgraded-dr.kanji/common"
)

func createFlexMessage(quiz common.Quiz) ([]byte, error) {
	// Use the format matched the quiz type
	var json []byte
	switch quiz.Type {
	case "Antonyms":
		json = common.AntonymsFormat
	case "Homonym":
		json = common.HomonymFormat
	case "Synonyms":
		json = common.SynonymsFormat
	case "Confer":
		json = common.ConferFormat
	case "Writing":
		json = common.WritingFormat
	case "Reading":
		json = common.ReadingFormat
	case "":
		return nil, errors.New("the quiz type is empty")
	default:
		return nil, errors.New("it is not supported the quiz type")
	}

	// var quizKanji string
	// if quizType[0] != "同音異義語" {
	// 	quizReg := regexp.MustCompile(`^(.*)（(.*)）$`)
	// 	quizKanji = string(quizReg.FindStringSubmatch(quiz[quizInIndex])[1])
	// } else {
	// 	if quizInIndex == 0 {
	// 		quizInIndex = 1
	// 	}
	// 	quizKanji = quiz[quizInIndex]
	// }

	// result = strings.Replace(
	// 	result,
	// 	"${quiz_kanji}",
	// 	quizKanji,
	// 	1,
	// )
	// result = strings.Replace(
	// 	result,
	// 	"${quiz_type}",
	// 	quizType[0],
	// 	1,
	// )
	// result = strings.Replace(
	// 	result,
	// 	"${id}",
	// 	fmt.Sprintf("%s %d", quizType[1], quizIndex),
	// 	2,
	// )

	return json, nil
}
