package answer

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Response(event *linebot.Event, message string) error {
	splitted := strings.Split(message, " ")

	switch splitted[1] {
	case "Antonym":
		err := antonym(event, splitted[2:])
		if err != nil {
			return err
		}

	case "Homonym":
		err := homonym(event, splitted[2:])
		if err != nil {
			return err
		}

	case "Synonym":
		err := synonym(event, splitted[2:])
		if err != nil {
			return err
		}

	case "Confer":
		err := confer(event, splitted[2:])
		if err != nil {
			return err
		}

	case "Writing":
		err := writing(event, splitted[2:])
		if err != nil {
			return err
		}

	case "Reading":
		err := reading(event, splitted[2:])
		if err != nil {
			return err
		}

	default:
		err := invalid(event)
		if err != nil {
			return err
		}
	}

	return nil
}
