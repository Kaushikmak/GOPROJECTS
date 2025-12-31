package add

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
)

func Add(taskDescription []string) {

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

	// get task as string
	taskDesc := strings.Join(taskDescription[2:], " ")
	if taskDesc == "" {
		fmt.Fprintf(os.Stderr, "Error!, trying to add empty task")
		return
	}

	task := models.Task{
		Id:          uuid.New(),
		Description: taskDesc,
		Status:      models.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	if err := fileio.Save(path, tasks); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)\n", len(tasks))
}
