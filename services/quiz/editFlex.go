package quiz

import (
	"fmt"
	"regexp"
	"strings"
)

func editFlex(base []byte) ([]byte, error) {
	var result string = string(base)

	quizType, quiz, quizIndex := device()

	quizReg := regexp.MustCompile(`^(.*)（(.*)）$`)
	// fmt.Println(quiz[0])
	// fmt.Println(quizReg.FindStringSubmatch(quiz[0]))
	// fmt.Println(len(quizReg.FindStringSubmatch(quiz[0])))
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
		2,
	)

	return []byte(result), nil
}
