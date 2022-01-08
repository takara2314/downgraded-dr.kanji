package state

import "time"

type State struct {
	LastReceive   time.Time
	RapidCount    int
	IsRapidNotice bool
	Quiz          struct {
		Enable bool
		Type   string
		No     int
		Option string
	}
}

var (
	States map[string]*State
)

func init() {
	States = map[string]*State{}
}
