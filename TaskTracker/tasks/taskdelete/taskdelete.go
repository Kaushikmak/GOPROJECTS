package taskdelete

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
)

func Delete(args []string) {
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Error: task id (or alias) required")
		return
	}

	idOrAlias := args[2]

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

	tasks = append(tasks[:index], tasks[index+1:]...)

	if err := fileio.Save(path, tasks); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Task deleted")
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
