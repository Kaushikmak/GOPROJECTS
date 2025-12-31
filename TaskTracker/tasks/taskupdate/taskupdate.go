package taskupdate

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
)

func Update(args []string) {
	if len(args) < 4 {
		fmt.Fprintln(os.Stderr, "Error: task id (or alias) and new description required")
		return
	}

	idOrAlias := args[2]
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

	index := findTaskIndex(tasks, idOrAlias)
	if index == -1 {
		fmt.Fprintf(os.Stderr, "Task not found: %s\n", idOrAlias)
		return
	}

	tasks[index].Description = newDesc
	tasks[index].UpdatedAt = time.Now()

	if err := fileio.Save(path, tasks); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Task updated:", tasks[index].Id)
}

// findTaskIndex handles both "1" (Alias) and "a1b2..." (UUID)
func findTaskIndex(tasks []models.Task, input string) int {
	// 1. Try to parse as Integer Alias (1, 2, 3...)
	if idx, err := strconv.Atoi(input); err == nil {
		realIndex := idx - 1 // Convert 1-based to 0-based
		if realIndex >= 0 && realIndex < len(tasks) {
			return realIndex
		}
	}

	// 2. Fallback: Search by UUID Prefix
	for i, t := range tasks {
		if strings.HasPrefix(t.Id.String(), input) {
			return i
		}
	}
	return -1
}
