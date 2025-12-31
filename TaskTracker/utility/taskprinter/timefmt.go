package taskprinter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
)

// stripANSI removes color codes to calculate the true visual length of a string.
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(str string) int {
	return len(ansiRegex.ReplaceAllString(str, ""))
}

// wordWrap splits a string into lines of a specific limit.
func wordWrap(text string, limit int) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	currentLine := words[0]

	for _, word := range words[1:] {
		if visibleLen(currentLine)+1+visibleLen(word) > limit {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			currentLine += " " + word
		}
	}
	lines = append(lines, currentLine)
	return lines
}

// PrintTasksColumn manually aligns columns to support colored output perfectly.
func PrintTasksColumn(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// 1. Calculate Column Widths
	// We set minimum widths for headers
	wKey := 3       // "#"
	wStatus := 6    // "STATUS"
	wCreated := 16  // "CREATED"
	wID := 9        // "ID (UUID)"
	
	for _, t := range tasks {
		// Check Key width
		if l := len(fmt.Sprintf("%d", t.Key)); l > wKey {
			wKey = l
		}
		// Check Status width (using visible length only)
		if l := visibleLen(t.Status.String()); l > wStatus {
			wStatus = l
		}
		// Check ID width (short ID is 8 chars)
		if 8 > wID {
			wID = 8
		}
	}

	// Add padding between columns
	gap := "   " 

	// 2. Print Header
	// We construct a format string based on calculated widths
	// %-*s means: Left-align (minus), width provided as arg (star), string (s)
	headerFmt := fmt.Sprintf("%%-%ds%%s%%-%ds%%s%%-%ds%%s%%-%ds%%s%%s\n", wKey, wStatus, wCreated, wID)
	fmt.Printf(headerFmt, "#", gap, "STATUS", gap, "CREATED", gap, "ID (UUID)", gap, "DESCRIPTION")

	// 3. Print Rows
	for _, t := range tasks {
		created := t.CreatedAt.Format("2006-01-02 15:04")
		statusStr := models.ColoredStatus(t.Status)
		shortID := t.Id.String()[:8]
		
		// Wrap Description
		// Dynamic wrap limit based on terminal size is complex, so we stick to a safe 50-60 chars
		descLines := wordWrap(t.Description, 60)

		// Calculate padding needed for the colored status
		// Because we print the raw ANSI string, we must add spaces manually based on visible length difference
		statusPadding := strings.Repeat(" ", wStatus-visibleLen(t.Status.String()))
		
		// Print First Line
		// We format parts individually to handle the color string + manual padding
		fmt.Printf("%-*d%s%s%s%s%-*s%s%-*s%s%s\n", 
			wKey, t.Key,           // Column 1: Key
			gap,                   // Gap
			statusStr, statusPadding, // Column 2: Status (Color + Manual Pad)
			gap,                   // Gap
			wCreated, created,     // Column 3: Created
			gap,                   // Gap
			wID, shortID,          // Column 4: ID
			gap,                   // Gap
			descLines[0],          // Column 5: Desc Line 1
		)

		// Print Continuation Lines
		if len(descLines) > 1 {
			// Create an empty prefix string that matches the width of the first 4 columns + gaps
			totalPrefixWidth := wKey + len(gap) + wStatus + len(gap) + wCreated + len(gap) + wID + len(gap)
			emptyPrefix := strings.Repeat(" ", totalPrefixWidth)

			for _, line := range descLines[1:] {
				fmt.Printf("%s%s\n", emptyPrefix, line)
			}
		}
	}
}
