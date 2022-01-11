package quiz

import (
	"errors"
	"fmt"
	"time"

	"downgraded-dr.kanji/common"
	"downgraded-dr.kanji/utils"
)

func choice(parameters []string) (common.Quiz, error) {
	// Debug: start time logging.
	startTime := time.Now()

	// Create a quiz instance.
	quiz := common.Quiz{}

	// Choice a quiz type if user does not select it.
	if parameters[1] == "" {
		quiz.Type = utils.RandChoiceString(common.QuizTypes)
	} else {
		quiz.Type = parameters[1]
	}

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
	default:
		return quiz, errors.New("invalid parameter")
	}

	fmt.Printf("[DEBUG] %s quizzing: %fms\n", quiz.Type, float64(time.Since(startTime)/time.Millisecond))

	return quiz, nil
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
