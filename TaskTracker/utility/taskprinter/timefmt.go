package taskprinter

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/term" // Standard package for terminal handling
	"github.com/kaushikmak/go-projects/TaskTracker/models"
)

// ansiRegex is used to strip color codes for accurate length calculation
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// visibleLen returns the length of the string without color codes
func visibleLen(str string) int {
	return len(ansiRegex.ReplaceAllString(str, ""))
}

// wordWrap splits a string into lines that fit within 'limit' characters.
// It ensures words are not broken in the middle; they are moved to the next line.
func wordWrap(text string, limit int) []string {
	if limit < 5 {
		limit = 5 // Safety net for extremely small screens
	}
	
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	currentLine := words[0]

	for _, word := range words[1:] {
		// Calculate length if we add this word (including a space)
		if visibleLen(currentLine)+1+visibleLen(word) > limit {
			// If it exceeds limit, push current line and start a new one
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			// Otherwise, append to current line
			currentLine += " " + word
		}
	}
	lines = append(lines, currentLine)
	return lines
}

// PrintTasksColumn prints the tasks in a table that adapts to terminal width.
func PrintTasksColumn(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// 1. Get Terminal Width
	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || termWidth <= 0 {
		termWidth = 80 // Fallback if not running in a real terminal
	}

	// 2. Calculate Column Widths (Dynamic based on content)
	wKey := 3      // "#"
	wStatus := 6   // "STATUS"
	wCreated := 16 // "CREATED"
	wID := 9       // "ID (UUID)"
	gapSize := 3   // Space between columns

	for _, t := range tasks {
		// Key width
		if l := len(fmt.Sprintf("%d", t.Key)); l > wKey {
			wKey = l
		}
		// Status width (visible length only)
		if l := visibleLen(t.Status.String()); l > wStatus {
			wStatus = l
		}
	}

	// Total width consumed by columns BEFORE Description
	// Key + gap + Status + gap + Created + gap + ID + gap
	prefixWidth := wKey + gapSize + wStatus + gapSize + wCreated + gapSize + wID + gapSize
	
	// Calculate available width for Description
	descWidth := termWidth - prefixWidth
	// Ensure a minimum width for readability (even if it breaks table alignment on tiny screens)
	if descWidth < 15 {
		descWidth = 15
	}

	// 3. Prepare Format Strings
	gap := strings.Repeat(" ", gapSize)
	headerFmt := fmt.Sprintf("%%-%ds%%s%%-%ds%%s%%-%ds%%s%%-%ds%%s%%s\n", wKey, wStatus, wCreated, wID)
	
	// Print Header
	fmt.Printf(headerFmt, "#", gap, "STATUS", gap, "CREATED", gap, "ID (UUID)", gap, "DESCRIPTION")

	// 4. Print Rows
	for _, t := range tasks {
		created := t.CreatedAt.Format("2006-01-02 15:04")
		statusStr := models.ColoredStatus(t.Status)
		shortID := t.Id.String()[:8]

		// Wrap the description based on the DYNAMIC available width
		descLines := wordWrap(t.Description, descWidth)

		// Calculate padding for status color correction
		statusPadding := strings.Repeat(" ", wStatus-visibleLen(t.Status.String()))

		// Print First Line
		fmt.Printf("%-*d%s%s%s%s%-*s%s%-*s%s%s\n", 
			wKey, t.Key,           
			gap,                   
			statusStr, statusPadding, 
			gap,                   
			wCreated, created,     
			gap,                   
			wID, shortID,          
			gap,                   
			descLines[0],          
		)

		// Print Continuation Lines (aligned under Description)
		if len(descLines) > 1 {
			// Create empty space matching the width of all previous columns
			emptyPrefix := strings.Repeat(" ", prefixWidth)
			for _, line := range descLines[1:] {
				fmt.Printf("%s%s\n", emptyPrefix, line)
			}
		}
	}
}
