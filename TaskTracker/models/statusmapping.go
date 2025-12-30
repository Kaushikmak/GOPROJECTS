package models

import "strings"

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
	switch strings.ToLower(s) {
	case "todo":
		return TODO
	case "in-progress", "inprogress":
		return IN_PROGRESS
	case "done":
		return DONE
	default:
		return UNKNOWN
	}
}


func (s TASKSTATUS) String() string {
	switch s {
	case TODO:
		return "TODO"
	case IN_PROGRESS:
		return "IN_PROGRESS"
	case DONE:
		return "DONE"
	default:
		return "UNKNOWN"
	}
}
