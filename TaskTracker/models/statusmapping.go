package models

func StatusToString(status TASKSTATUS) string {
	switch status {
	case TODO:
		return "TODO"
	case IN_PROGRESS:
		return "IN-PROGRESS"
	case DONE:
		return "DONE"
	default:
		return "UNKNOWN"
	}
}

func StringToStatus(s string) TASKSTATUS {
	switch s {
	case "TODO":
		return TODO
	case "IN-PROGRESS":
		return IN_PROGRESS
	case "DONE":
		return DONE
	default:
		return UNKNOWN
	}
}
