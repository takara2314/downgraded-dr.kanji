package quiz

import (
	"strconv"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func writing(quiz *common.Quiz) {
	// Choice a section
	quiz.Section = utils.RandChoiceString(common.QuizSections)

	switch quiz.Section {
	case "antonyms":
		writingFromAntonyms(quiz)
	case "homonyms":
		writingFromHomonyms(quiz)
	case "synonyms":
		writingFromSynonyms(quiz)
	case "confers":
		writingFromConfers(quiz)
	case "others":
		writingFromOthers(quiz)
	}
}

func writingFromAntonyms(quiz *common.Quiz) {
	// Choice a quiz no
	quiz.No = utils.RandN(len(common.Quizzes.Antonyms)) + 1

	// Choice a quiz content no
	contentNo := utils.RandN(len(common.Quizzes.Antonyms[quiz.No-1]))

	// Make a option
	quiz.Option = strconv.Itoa(contentNo + 1)

	// Make a content
	quiz.Content = utils.HiraganaToKatakana(
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Antonyms[quiz.No-1][contentNo]),
		),
	)
}

func writingFromHomonyms(quiz *common.Quiz) {
	// Choice a quiz no
	quiz.No = utils.RandN(len(common.Quizzes.Homonyms)) + 1

	// Make a option
	quiz.Option = "1"

	// Make a content
	quiz.Content = utils.HiraganaToKatakana(
		snipOtherMoji(
			common.Quizzes.Homonyms[quiz.No-1][0],
		),
	)
}

func writingFromSynonyms(quiz *common.Quiz) {
	// Choice a quiz no
	quiz.No = utils.RandN(len(common.Quizzes.Synonyms)) + 1

	// Choice a quiz content no
	contentNo := utils.RandN(len(common.Quizzes.Synonyms[quiz.No-1]))

	// Make a option
	quiz.Option = strconv.Itoa(contentNo + 1)

	// Make a content
	quiz.Content = utils.HiraganaToKatakana(
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Synonyms[quiz.No-1][contentNo]),
		),
	)
}

func writingFromConfers(quiz *common.Quiz) {
	// Choice a quiz no
	quiz.No = utils.RandN(len(common.Quizzes.Confers)) + 1

	// Choice a quiz content no
	contentNo := utils.RandN(len(common.Quizzes.Confers[quiz.No-1]))

	// Make a option
	quiz.Option = strconv.Itoa(contentNo + 1)

	// Make a content
	quiz.Content = utils.HiraganaToKatakana(
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Confers[quiz.No-1][contentNo]),
		),
	)
}

func writingFromOthers(quiz *common.Quiz) {
	// Choice a quiz no
	quiz.No = utils.RandN(len(common.Quizzes.Others)) + 1

	// Make a option
	quiz.Option = "1"

	// Make a content
	quiz.Content = utils.HiraganaToKatakana(
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Others[quiz.No-1]),
		),
	)
}
