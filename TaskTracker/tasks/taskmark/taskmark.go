package taskmark

import (
	"fmt"
	"os"
	"time"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
)


func Mark(args []string) {
	if len(args) < 4 {
		fmt.Fprintln(os.Stderr, "Usage: task-cli mark <id> <status>")
		return
	}

	id := args[2]

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

	index := findTaskIndex(tasks, id)
	if index == -1 {
		fmt.Fprintf(os.Stderr, "Task not found: %s\n", id)
		return
	}

	// mutate
	tasks[index].Status = status
	tasks[index].UpdatedAt = time.Now()

	if err := fileio.Save(path, tasks); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Task %s marked as %s\n", id, status.String())
}


func findTaskIndex(tasks []models.Task, shortID string) int {
	for i, t := range tasks {
		if len(t.Id.String()) >= 8 && t.Id.String()[:8] == shortID {
			return i
		}
	}
	return -1
}

