package state

import (
	"time"

	"downgraded-dr.kanji/common"
)

type State struct {
	LastReceive   time.Time
	RapidCount    int
	IsRapidNotice bool
	IsQuizzing    bool
	LastQuiz      common.Quiz
}

var (
	States map[string]*State
)

func init() {
	States = map[string]*State{}
}
