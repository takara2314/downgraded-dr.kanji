package quiz

import (
	"math/rand"
	"time"

	"downgraded-dr.kanji/common"
)

func device() ([]string, []string, int, int) {
	// シード更新
	rand.Seed(time.Now().Unix())

	types := [][]string{
		{"対義語", "Antonyms"},
		{"同音異義語", "Homonym"},
		{"類義語", "Synonyms"},
		{"参照(cf)", "Confer"},
		{"三字熟語", "Three"},
		{"四字熟語", "Four"},
	}

	// 出題タイプ (index)
	quizTypeIndex := rand.Intn(len(types))

	// 出題タイプの問題集
	var quizes [][]string
	switch quizTypeIndex {
	case 0:
		quizes = common.Quizzes.Antonyms
	case 1:
		quizes = common.Quizzes.Homonym
	case 2:
		quizes = common.Quizzes.Synonyms
	case 3:
		quizes = common.Quizzes.Confer
	case 4:
		quizes = common.Quizzes.Three
	case 5:
		quizes = common.Quizzes.Four
	}

	// 問題 (index)
	quizIndex := rand.Intn(len(quizes))

	// 問題の何要素目
	quizInIndex := rand.Intn(len(quizes[quizIndex]))
	// 四字熟語なら必然的に0
	if quizTypeIndex == 5 {
		quizInIndex = 0
	}

	return types[quizTypeIndex], quizes[quizIndex], quizIndex, quizInIndex
}
