package tasklist

import (
	"fmt"
	"os"
	"strings"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/tasks/help"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/taskprinter"
)

func filternPrint(tasks []models.Task, status models.TASKSTATUS) {
	var filtered []models.Task

	for _, t := range tasks {
		if t.Status == status {
			filtered = append(filtered, t)
		}
	}

	taskprinter.PrintTasksColumn(filtered)
}

func List(taskDescription []string) {
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

	// ASSIGN ALIAS KEYS (1-based index)
	for i := range tasks {
		tasks[i].Key = i + 1
	}

	taskDesc := strings.Join(taskDescription[2:], " ")

	if taskDesc == "" {
		taskprinter.PrintTasksColumn(tasks)
	} else {
		switch models.StringToStatus(taskDesc) {
		case models.DONE:
			filternPrint(tasks, models.DONE)
		case models.IN_PROGRESS:
			filternPrint(tasks, models.IN_PROGRESS)
		case models.TODO:
			filternPrint(tasks, models.TODO)
		case models.UNKNOWN:
			fmt.Printf("Invalid command: %s\n", strings.ToLower(taskDesc))
			help.ShowOptions()
			return
		default:
			fmt.Println("Invalid command")
			help.ShowOptions()
			return
		}
	}
}
