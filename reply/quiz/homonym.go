package quiz

import (
	"fmt"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func homonym(quiz *common.Quiz) {
	// Set a section.
	quiz.Section = "homonyms"

	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Homonyms)) + 1

	// Choice a quiz content no.
	contentNo := utils.RandMN(1, len(common.Quizzes.Homonyms[quiz.No-1]))

	// Make a quiz content.
	quiz.Content = snipOtherMoji(
		common.Quizzes.Homonyms[quiz.No-1][contentNo],
	)

	quiz.Option = fmt.Sprintf("%d", contentNo)

	// Make a suggested correct answer.
	corrects := utils.StringSliceRemove(common.Quizzes.Homonyms[quiz.No-1], contentNo)
	quiz.Corrects = utils.StringSliceRemove(corrects, 0)
}
