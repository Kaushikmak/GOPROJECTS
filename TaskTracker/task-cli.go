package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/kaushikmak/go-projects/TaskTracker/tasks/add"
	"github.com/kaushikmak/go-projects/TaskTracker/tasks/help"
	"github.com/kaushikmak/go-projects/TaskTracker/tasks/taskdelete"
	"github.com/kaushikmak/go-projects/TaskTracker/tasks/tasklist"
	"github.com/kaushikmak/go-projects/TaskTracker/tasks/taskmark"
	"github.com/kaushikmak/go-projects/TaskTracker/tasks/taskupdate"
)

var COMMANDS = []string{"add", "update", "delete", "mark", "list", "help"}

func main() {
	// fetch command given by user
	args := os.Args
	// throw error if now command is given by user and show help
	if len(args) <= 1 {
		fmt.Fprintf(os.Stderr, "Error no command given by user\n")
		fmt.Println("for help type ")
		return
	}
	cmd := args[1]
	// if invalid command
	if !slices.Contains(COMMANDS, cmd) {
		fmt.Fprintf(os.Stderr, "Error: '%s' is not a valid command.\n", cmd)
		help.ShowOptions()
        return
	}

	switch cmd {
	case "help":
		help.ShowOptions()
	case "add":
		add.Add(args)
    case "list":
       tasklist.List(args) 
   case "delete":
        taskdelete.Delete(args)
    case "update":
        taskupdate.Update(args)
    case "mark":
        taskmark.Mark(args)
    default:
        help.ShowOptions()
        return
    }

}
