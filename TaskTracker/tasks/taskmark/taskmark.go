package taskmark

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
)

func Mark(args []string) {
	if len(args) < 4 {
		fmt.Fprintln(os.Stderr, "Usage: task-cli mark <id|alias> <status>")
		return
	}

	idOrAlias := args[2]

	status := models.StringToStatus(args[3])
	if status == models.UNKNOWN {
		fmt.Fprintf(os.Stderr, "Invalid status: %s\n", args[3])
		return
	}

	path, err := fileio.EnsureStorage()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	tasks, err := fileio.Load(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	index := findTaskIndex(tasks, idOrAlias)
	if index == -1 {
		fmt.Fprintf(os.Stderr, "Task not found: %s\n", idOrAlias)
		return
	}

	// mutate
	tasks[index].Status = status
	tasks[index].UpdatedAt = time.Now()

	if err := fileio.Save(path, tasks); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// NEW: Confirmation
	fmt.Printf("Task %s marked as %s\n", idOrAlias, status.String())
}

func findTaskIndex(tasks []models.Task, input string) int {
	if idx, err := strconv.Atoi(input); err == nil {
		realIndex := idx - 1
		if realIndex >= 0 && realIndex < len(tasks) {
			return realIndex
		}
	}
	for i, t := range tasks {
		if strings.HasPrefix(t.Id.String(), input) {
			return i
		}
	}
	return -1
}
