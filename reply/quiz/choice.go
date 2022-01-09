package quiz

import (
	"fmt"
	"time"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func choice() common.Quiz {
	// Debug: start time logging.
	startTime := time.Now()

	// Create a quiz instance.
	quiz := common.Quiz{}

	// Choice a quiz type.
	quiz.Type = utils.RandChoiceString(common.QuizTypes)

	switch quiz.Type {
	case "Antonym":
		antonym(&quiz)
	case "Homonym":
		homonym(&quiz)
	case "Synonym":
		synonym(&quiz)
	case "Confer":
		confer(&quiz)
	case "Writing":
		writing(&quiz)
	case "Reading":
		reading(&quiz)
	}

	fmt.Printf("[DEBUG] %s quizzing: %fms\n", quiz.Type, float64(time.Since(startTime)/time.Millisecond))

	return quiz
}

// snipOtherMoji returns the moji snipped other moji.
func snipOtherMoji(s string) string {
	return utils.SnipStringCovered(s, "[", "]")
}

// snipYomiMoji returns the moji snipped yomi.
func snipYomiMoji(s string) string {
	return utils.SnipStringCovered(s, "（", "）")
}

// obtainYomiMoji returns the yomi only.
func obtainYomiMoji(s string) string {
	return utils.ObtainStringCovered(s, "（", "）")
}
