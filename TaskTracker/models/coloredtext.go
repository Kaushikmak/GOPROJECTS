package models

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

func ColoredStatus(s TASKSTATUS) string {
	switch s {
	case TODO:
		return colorYellow + s.String() + colorReset
	case IN_PROGRESS:
		return colorRed + s.String() + colorReset
	case DONE:
		return colorGreen + s.String() + colorReset
	default:
		return s.String()
	}
}
