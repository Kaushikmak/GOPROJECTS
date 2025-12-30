package taskdelete

import (
	"fmt"
	"os"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
)


func Delete(args []string) {
	// validate args
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Error: task id required")
		return
	}

	id := args[2]

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

	// remove task
	tasks = append(tasks[:index], tasks[index+1:]...)

	if err := fileio.Save(path, tasks); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Task deleted:", id)
}


func findTaskIndex(tasks []models.Task, shortID string) int {
	for i, t := range tasks {
		if len(t.Id.String()) >= 8 && t.Id.String()[:8] == shortID {
			return i
		}
	}
	return -1
}

