package quiz

import (
	"fmt"
	"strconv"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func choice() common.Quiz {
	// Create a quiz instance
	quiz := common.Quiz{}

	// Choice a quiz type
	quiz.Type = utils.RandChoiceString(common.QuizTypes)

	switch quiz.Type {
	case "Antonyms":
		// Choice a quiz no
		quiz.No = utils.RandN(len(common.Quizzes.Antonyms))

		// Base of a quiz content
		var contents [2]string
		contents[0] = snipOtherMoji(
			snipYomiMoji(common.Quizzes.Antonyms[quiz.No][0]),
		)
		contents[1] = snipOtherMoji(
			snipYomiMoji(common.Quizzes.Antonyms[quiz.No][1]),
		)

		// Choice blank content side
		side := utils.RandN(2)
		// Choice blank position
		blank := utils.RandN(utils.LenString(contents[side]))
		quiz.Option = strconv.Itoa(blank + 1)

		// Make blank
		contents[side] = utils.ReplaceStringFromIndex(contents[side], "□", blank)

		// Concat contents
		quiz.Content = fmt.Sprintf("%s ←→ %s", contents[0], contents[1])
	}

	return quiz

	// // 出題タイプ (index)
	// quizTypeIndex := rand.Intn(len(types))

	// // 出題タイプの問題集
	// var quizes [][]string
	// switch quizTypeIndex {
	// case 0:
	// 	quizes = common.Quizzes.Antonyms
	// case 1:
	// 	quizes = common.Quizzes.Homonym
	// case 2:
	// 	quizes = common.Quizzes.Synonyms
	// case 3:
	// 	quizes = common.Quizzes.Confer
	// case 4:
	// 	quizes = common.Quizzes.Three
	// case 5:
	// 	quizes = common.Quizzes.Four
	// }

	// // 問題 (index)
	// quizIndex := rand.Intn(len(quizes))

	// // 問題の何要素目
	// quizInIndex := rand.Intn(len(quizes[quizIndex]))
	// // 四字熟語なら必然的に0
	// if quizTypeIndex == 5 {
	// 	quizInIndex = 0
	// }

	// return types[quizTypeIndex], quizes[quizIndex], quizIndex, quizInIndex
}

// snipOtherMoji returns the moji snipped other moji.
func snipOtherMoji(s string) string {
	return utils.SnipStringCovered(s, "[", "]")
}

// snipYomiMoji returns the moji snipped yomi.
func snipYomiMoji(s string) string {
	return utils.SnipStringCovered(s, "（", "）")
}
