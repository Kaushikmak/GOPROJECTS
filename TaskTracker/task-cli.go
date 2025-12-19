package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/kaushikmak/go-projects/TaskTracker/tasks/add"
	"github.com/kaushikmak/go-projects/TaskTracker/tasks/help"
)

var COMMANDS = []string{"add", "update", "delete", "mark-in-progress", "mark-done", "list", "help"}

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
	}

	switch cmd {
	case "help":
		help.ShowOptions()
	case "add":
		add.Add(args)

	}

}
