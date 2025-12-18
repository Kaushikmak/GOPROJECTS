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

	fmt.Println(task)

	x, err := fileio.ReadData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading")
	}
	fmt.Println(string(x))

}
