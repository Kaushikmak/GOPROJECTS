package taskprinter

import (
	"fmt"
	"time"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
)

/*
Column widths (single source of truth)
*/
const (
	colIndexWidth  = 3
	colIDWidth     = 8
	colDescWidth   = 30
	colStatusWidth = 14
	colTimeWidth   = 22
)

/*
Human-friendly time formatting
*/
func HumanTime(t time.Time) string {
	now := time.Now()

	// Normalize to midnight
	y, m, d := now.Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, now.Location())

	ty, tm, td := t.Date()
	thatDay := time.Date(ty, tm, td, 0, 0, 0, 0, t.Location())

	diffDays := int(thatDay.Sub(today).Hours() / 24)

	switch diffDays {
	case 0:
		return "Today " + t.Format("3:04 PM")
	case -1:
		return "Yesterday " + t.Format("3:04 PM")
	case 1:
		return "Tomorrow " + t.Format("3:04 PM")
	default:
		return t.Format("Mon, 02 Jan 2006 3:04 PM")
	}
}

/*
Main printer
*/
func PrintTasksColumn(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// Header
	fmt.Printf(
		"%-3s %-8s %-30s %-14s %-22s %-22s\n",
		"#",
		"ID",
		"TASK DESCRIPTION",
		"STATUS",
		"UPDATED",
		"CREATED",
	)

	for i, t := range tasks {
		printWrappedTask(i+1, t)
	}
}

/*
Print one task with wrapped description
*/
func printWrappedTask(index int, t models.Task) {
	descLines := wrapText(t.Description, colDescWidth)

	for lineIdx, line := range descLines {
		if lineIdx == 0 {
			// First line: full row
			fmt.Printf(
				"%-3d %-8s %-30s %-14s %-22s %-22s\n",
				index,
				shortID(t.Id.String()),
				line,
				models.ColoredStatus(t.Status),
				HumanTime(t.UpdatedAt),
				HumanTime(t.CreatedAt),
			)
		} else {
			// Wrapped lines: indent other columns
			fmt.Printf(
				"%-3s %-8s %-30s %-14s %-22s %-22s\n",
				"",
				"",
				line,
				"",
				"",
				"",
			)
		}
	}
}

/*
Helpers
*/
func shortID(id string) string {
	if len(id) > 8 {
		return id[:8]
	}
	return id
}

func wrapText(s string, width int) []string {
	var lines []string

	for len(s) > width {
		lines = append(lines, s[:width])
		s = s[width:]
	}

	if len(s) > 0 {
		lines = append(lines, s)
	}

	return lines
}

