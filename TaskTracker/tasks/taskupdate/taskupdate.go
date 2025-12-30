package taskupdate

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
)


func Update(args []string) {
	// validate args
	if len(args) < 4 {
		fmt.Fprintln(os.Stderr, "Error: task id and new description required")
		return
	}

	id := args[2]
	newDesc := strings.Join(args[3:], " ")
	if newDesc == "" {
		fmt.Fprintln(os.Stderr, "Error: empty description")
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

	// update task
	tasks[index].Description = newDesc
	tasks[index].UpdatedAt = time.Now()

	if err := fileio.Save(path, tasks); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Task updated:", id)
}


func findTaskIndex(tasks []models.Task, shortID string) int {
	for i, t := range tasks {
		if len(t.Id.String()) >= 8 && t.Id.String()[:8] == shortID {
			return i
		}
	}
	return -1
}

