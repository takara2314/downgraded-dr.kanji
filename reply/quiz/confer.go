package quiz

import (
	"fmt"
	"sort"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func confer(quiz *common.Quiz) {
	// Choice a quiz no
	quiz.No = utils.RandN(len(common.Quizzes.Confers))

	// Choice a quiz content no
	contentNo := utils.RandMN(1, len(common.Quizzes.Homonyms[quiz.No]))

	// Base of a quiz content
	content := snipOtherMoji(
		snipYomiMoji(common.Quizzes.Antonyms[quiz.No][contentNo]),
	)

	// Choice blank positions
	var blanks []int = make([]int, 2)
	blanks[0] = utils.RandN(utils.LenString(content))
	blanks[1] = utils.RandN(utils.LenString(content))
	for {
		if blanks[0] != blanks[1] {
			break
		}
		blanks[1] = utils.RandN(utils.LenString(content))
	}

	sort.Ints(blanks)

	quiz.Option = fmt.Sprintf("%d_%d,%d", contentNo, blanks[0]+1, blanks[1]+1)

	// Make blank
	for _, blank := range blanks {
		content = utils.ReplaceStringFromIndex(content, "□", blank)
	}

	// Make a content
	quiz.Content = content
}
