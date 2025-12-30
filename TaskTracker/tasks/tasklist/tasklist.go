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

func filternPrint(tasks []models.Task,status models.TASKSTATUS) {
    var filtered []models.Task

    for _,t := range tasks{
        if t.Status == status {
            filtered = append(filtered, t)
        }
    }
   
    taskprinter.PrintTasksColumn(filtered)        
     
}

func listAll(tasks []models.Task) {
	taskprinter.PrintTasksColumn(tasks)
}

func listDone(tasks []models.Task) {
	filternPrint(tasks, models.DONE)
}

func listInProgress(tasks []models.Task) {
	filternPrint(tasks, models.IN_PROGRESS)
}

func listTodo(tasks []models.Task) {
	filternPrint(tasks, models.TODO)
}



func List(taskDescription []string){
    path, err := fileio.EnsureStorage()
    if err != nil {
        fmt.Fprintln(os.Stderr,err)
        return
    }

    tasks, err := fileio.Load(path)
    if err != nil {
        fmt.Fprintln(os.Stderr,err)
        return
    }

    taskDesc := strings.Join(taskDescription[2:]," ")
    
    if taskDesc == ""{
        listAll(tasks)
    }else{


    switch models.StringToStatus(taskDesc){
    case models.DONE:
        listDone(tasks)
    case models.IN_PROGRESS:
        listInProgress(tasks)
    case models.TODO:
        listTodo(tasks)
    case models.UNKNOWN:
        fmt.Printf("Invalid command: %s\n",strings.ToLower(taskDesc))
        help.ShowOptions()
        return
    default:
        fmt.Println("Invalid command")
        help.ShowOptions()
        return
    }


}
}
