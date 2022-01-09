package quiz

import (
	"fmt"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func synonym(quiz *common.Quiz) {
	// Set a section.
	quiz.Section = "synonyms"

	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Synonyms)) + 1

	// Choice a quiz content no.
	contentNo1 := utils.RandN(len(common.Quizzes.Synonyms[quiz.No-1]))
	contentNo2 := utils.RandN(len(common.Quizzes.Synonyms[quiz.No-1]))
	for {
		if contentNo1 != contentNo2 {
			break
		}
		contentNo2 = utils.RandN(len(common.Quizzes.Synonyms[quiz.No-1]))
	}

	// Base of a quiz content.
	var contents [2]string
	contents[0] = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Synonyms[quiz.No-1][contentNo1]),
	)
	contents[1] = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Synonyms[quiz.No-1][contentNo2]),
	)

	// Choice blank content side.
	side := utils.RandN(2)
	sideName := "L"
	if side == 1 {
		sideName = "R"
	}

	// Choice blank position.
	blank := utils.RandN(utils.LenString(contents[side]))

	quiz.Option = fmt.Sprintf("%d,%d_%s%d", contentNo1+1, contentNo2+1, sideName, blank+1)

	// Make blank.
	correct := utils.SliceString(contents[side], blank, blank+1)
	contents[side] = utils.ReplaceStringFromIndex(contents[side], "□", blank)

	// Concat contents.
	quiz.Content = fmt.Sprintf("%s ≒ %s", contents[0], contents[1])

	// Make a suggested correct answer.
	quiz.Corrects = []string{
		correct,
	}
}
