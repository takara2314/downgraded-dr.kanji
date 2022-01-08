package ask

type Ask struct {
	TargetId string
	Quiz     struct {
		Type   string
		No     int
		Option string
	}
}
