package help

import "fmt"

// shows all options and how to use task-cli
func ShowOptions() {
	fmt.Println("Following are valid commands")
	fmt.Println("help\t\t\tto show help menu")
	fmt.Println()

	fmt.Println("add\t\t\tto add new task")
	fmt.Println("\t\t\tusage:\n\t\t\t\ttask-cli add 'uninstall windows'")
	fmt.Println()

	fmt.Println("update\t\t\tto update previous task")
	fmt.Println("\t\t\tusage:\n\t\t\t\ttask-cli update 1 'uninstall windows and install arch'")
	fmt.Println()

	fmt.Println("delete\t\t\tto delete existing task")
	fmt.Println("\t\t\tusage:\n\t\t\t\ttask-cli delete 1")
	fmt.Println()

	fmt.Println("mark-in-progress\tto mark task as in-progress")
	fmt.Println("\t\t\tusage:\n\t\t\t\ttask-cli mark-in-progress 1")
	fmt.Println()

	fmt.Println("mark-done\t\tto mark task as done")
	fmt.Println("\t\t\tusage:\n\t\t\t\ttask-cli mark-done 1")
	fmt.Println()

	fmt.Println("list\t\t\tto list all task")
	fmt.Println("\t\t\tusage:\n\t\t\t\ttask-cli list")
	fmt.Println("\t\t\t\ttask-cli list done")
	fmt.Println("\t\t\t\ttask-cli list todo")
	fmt.Println("\t\t\t\ttask-cli list in-progress")
}
