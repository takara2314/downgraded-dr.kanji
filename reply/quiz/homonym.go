package quiz

import (
	"fmt"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func homonym(quiz *common.Quiz) {
	// Choice a quiz no
	quiz.No = utils.RandN(len(common.Quizzes.Homonyms))

	// Choice a quiz content no
	contentNo := utils.RandMN(1, len(common.Quizzes.Homonyms[quiz.No]))

	// Make a quiz content
	quiz.Content = snipOtherMoji(
		common.Quizzes.Homonyms[quiz.No][contentNo],
	)

	quiz.Option = fmt.Sprintf("%d", contentNo)
}
