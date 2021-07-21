package quiz

import (
	"fmt"
	"regexp"
	"strings"
)

func editFlex(base []byte) ([]byte, error) {
	var result string = string(base)

	quizType, quiz, quizIndex := device()

	quizReg := regexp.MustCompile(`^([0-9]|[０-９]){5}$`)
	quizKanji := string(quizReg.FindStringSubmatch(quiz[0])[1])

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
		1,
	)

	return []byte(result), nil
}
