package taskprinter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
)

// PrintTasksColumn displays tasks in a perfectly aligned table.
func PrintTasksColumn(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// Initialize tabwriter
	// minwidth: 0 (minimal cell width)
	// tabwidth: 0 (tab width)
	// padding:  3 (spaces between columns)
	// padchar:  ' ' (character to pad with)
	// flags:    0
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Header
	// We use \t (tab) to separate columns. tabwriter replaces these with the correct padding.
	fmt.Fprintln(w, "ID\tSTATUS\tCREATED\tDESCRIPTION")

	for _, t := range tasks {
		// Format Time (Format: YYYY-MM-DD HH:MM)
		created := t.CreatedAt.Format("2006-01-02 15:04")

		// Get Colored Status
		// Note: tabwriter handles ANSI color codes well as long as the escape sequence lengths are consistent.
		status := models.ColoredStatus(t.Status)

		// Print row
		// %v for UUID works automatically
		fmt.Fprintf(w, "%v\t%s\t%s\t%s\n", t.Id, status, created, t.Description)
	}

	// Important: Flush the writer to output the data to stdout
	w.Flush()
}
