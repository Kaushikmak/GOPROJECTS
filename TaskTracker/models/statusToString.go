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
