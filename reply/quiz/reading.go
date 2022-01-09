package quiz

import (
	"strconv"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func reading(quiz *common.Quiz) {
	// Choice a section.
	quiz.Section = utils.RandChoiceString(common.QuizSections)

	switch quiz.Section {
	case "antonyms":
		readingFromAntonyms(quiz)
	case "homonyms":
		readingFromHomonyms(quiz)
	case "synonyms":
		readingFromSynonyms(quiz)
	case "confers":
		readingFromConfers(quiz)
	case "others":
		readingFromOthers(quiz)
	}
}

func readingFromAntonyms(quiz *common.Quiz) {
	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Antonyms)) + 1

	// Choice a quiz content no.
	contentNo := utils.RandN(len(common.Quizzes.Antonyms[quiz.No-1]))

	// Make a option.
	quiz.Option = strconv.Itoa(contentNo + 1)

	// Make a content.
	quiz.Content = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Antonyms[quiz.No-1][contentNo]),
	)

	// Make a suggested correct answer.
	quiz.Corrects = []string{
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Antonyms[quiz.No-1][contentNo]),
		),
	}
}

func readingFromHomonyms(quiz *common.Quiz) {
	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Homonyms)) + 1

	// Choice a quiz content no.
	contentNo := utils.RandMN(1, len(common.Quizzes.Homonyms[quiz.No-1]))

	// Make a option.
	quiz.Option = strconv.Itoa(contentNo)

	// Make a content.
	quiz.Content = snipOtherMoji(
		common.Quizzes.Homonyms[quiz.No-1][contentNo],
	)

	// Make a suggested correct answer.
	quiz.Corrects = []string{
		snipOtherMoji(
			common.Quizzes.Homonyms[quiz.No-1][0],
		),
	}
}

func readingFromSynonyms(quiz *common.Quiz) {
	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Synonyms)) + 1

	// Choice a quiz content no.
	contentNo := utils.RandN(len(common.Quizzes.Synonyms[quiz.No-1]))

	// Make a option.
	quiz.Option = strconv.Itoa(contentNo + 1)

	// Make a content.
	quiz.Content = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Synonyms[quiz.No-1][contentNo]),
	)

	// Make a suggested correct answer.
	quiz.Corrects = []string{
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Synonyms[quiz.No-1][contentNo]),
		),
	}
}

func readingFromConfers(quiz *common.Quiz) {
	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Confers)) + 1

	// Choice a quiz content no.
	contentNo := utils.RandN(len(common.Quizzes.Confers[quiz.No-1]))

	// Make a option.
	quiz.Option = strconv.Itoa(contentNo + 1)

	// Make a content.
	quiz.Content = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Confers[quiz.No-1][contentNo]),
	)

	// Make a suggested correct answer.
	quiz.Corrects = []string{
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Confers[quiz.No-1][contentNo]),
		),
	}
}

func readingFromOthers(quiz *common.Quiz) {
	// Choice a quiz no.
	quiz.No = utils.RandN(len(common.Quizzes.Others)) + 1

	// Make a option.
	quiz.Option = "1"

	// Make a content.
	quiz.Content = snipOtherMoji(
		snipYomiMoji(common.Quizzes.Others[quiz.No-1]),
	)

	// Make a suggested correct answer.
	quiz.Corrects = []string{
		snipOtherMoji(
			obtainYomiMoji(common.Quizzes.Others[quiz.No-1]),
		),
	}
}
