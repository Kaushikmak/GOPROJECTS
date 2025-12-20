package add

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kaushikmak/go-projects/TaskTracker/models"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/fileio"
	"github.com/kaushikmak/go-projects/TaskTracker/utility/serialization"
)

func Add(taskDescription []string) {
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

	// fmt.Println(task)

	serializedTask := serialization.Serialize(task)

	err := fileio.AppendTask(serializedTask)
	if err != nil {
		fmt.Println("Error saving task: %v\n", err)
		return
	}
	fmt.Println("Task added successfully (ID:%s)\n", task.Id)

}
