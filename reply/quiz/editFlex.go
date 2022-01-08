package quiz

import (
	"fmt"
	"regexp"
	"strings"
)

func editFlex(base []byte) ([]byte, error) {
	var result string = string(base)

	quizType, quiz, quizIndex, quizInIndex := device()

	var quizKanji string
	if quizType[0] != "同音異義語" {
		quizReg := regexp.MustCompile(`^(.*)（(.*)）$`)
		quizKanji = string(quizReg.FindStringSubmatch(quiz[quizInIndex])[1])
	} else {
		if quizInIndex == 0 {
			quizInIndex = 1
		}
		quizKanji = quiz[quizInIndex]
	}

	result = strings.Replace(
		result,
		"${quiz_kanji}",
		quizKanji,
		1,
	)
	result = strings.Replace(
		result,
		"${quiz_type}",
		quizType[0],
		1,
	)
	result = strings.Replace(
		result,
		"${id}",
		fmt.Sprintf("%s %d", quizType[1], quizIndex),
		2,
	)

	return []byte(result), nil
}
