package quiz

import (
	"math/rand"
	"time"
)

func device() ([]string, []string, int, int) {
	// シード更新
	rand.Seed(time.Now().Unix())

	types := [][]string{
		{"対義語", "Antonyms"},
		{"同音異義語", "Homonym"},
		{"類義語", "Synonyms"},
		{"参照(cf)", "Confer"},
		{"四字熟語", "Four"},
	}

	// 出題タイプ (index)
	quizTypeIndex := rand.Intn(len(types))

	// 出題タイプの問題集
	var quizes [][]string
	switch quizTypeIndex {
	case 0:
		quizes = Config.Antonyms
	case 1:
		quizes = Config.Homonym
	case 2:
		quizes = Config.Synonyms
	case 3:
		quizes = Config.Confer
	case 4:
		quizes = Config.Four
	}

	// 問題 (index)
	quizIndex := rand.Intn(len(quizes))

	// 問題の何要素目
	quizInIndex := rand.Intn(len(quizes[quizIndex]))

	return types[quizTypeIndex], quizes[quizIndex], quizIndex, quizInIndex
}
