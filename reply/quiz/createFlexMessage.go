package quiz

import (
	"errors"
	"strconv"
	"strings"

	"downgraded-dr.kanji/common"
)

func createFlexMessage(quiz common.Quiz) ([]byte, error) {
	// Use the format matched the quiz type
	var json string
	switch quiz.Type {
	case "Antonym":
		json = string(common.AntonymFormat)
	case "Homonym":
		json = string(common.HomonymFormat)
	case "Synonym":
		json = string(common.SynonymFormat)
	case "Confer":
		json = string(common.ConferFormat)
	case "Writing":
		json = string(common.WritingFormat)
	case "Reading":
		json = string(common.ReadingFormat)
	case "":
		return nil, errors.New("the quiz type is empty")
	default:
		return nil, errors.New("it is not supported the quiz type")
	}

	// Insert quiz details
	json = strings.Replace(
		json,
		"${content}",
		quiz.Content,
		1,
	)
	json = strings.Replace(
		json,
		"${no}",
		strconv.Itoa(quiz.No),
		2,
	)
	json = strings.Replace(
		json,
		"${option}",
		quiz.Option,
		2,
	)
	json = strings.Replace(
		json,
		"${memo}",
		quiz.Memo,
		1,
	)
	json = strings.Replace(
		json,
		"${section}",
		quiz.Section,
		2,
	)

	return []byte(json), nil
}
