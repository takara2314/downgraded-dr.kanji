package quiz

import (
	"fmt"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func antonym(quiz *common.Quiz) {
	// Set a section.
	quiz.Section = "antonyms"

	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Antonyms)) + 1

	// Base of a quiz content.
	var contents [2]string
	contents[0] = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Antonyms[quiz.No-1][0]),
	)
	contents[1] = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Antonyms[quiz.No-1][1]),
	)

	// Choice blank content side.
	side := utils.RandN(2)
	sideName := "L"
	if side == 1 {
		sideName = "R"
	}

	// Choice blank position.
	blank := utils.RandN(utils.LenString(contents[side]))

	quiz.Option = fmt.Sprintf("%s%d", sideName, blank+1)

	// Make blank.
	contents[side] = utils.ReplaceStringFromIndex(contents[side], "□", blank)

	// Concat contents.
	quiz.Content = fmt.Sprintf("%s ←→ %s", contents[0], contents[1])
}
