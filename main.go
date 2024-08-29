package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'update', 'delete', 'list', 'done', 'notdone', or 'inprogress' command.")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask(os.Args[2:])
	case "update":
		updateTask(os.Args[2:])
	case "delete":
		deleteTask(os.Args[2:])
	case "list":
		listTasks()
	case "done":
		listTasksByStatus("done")
	case "notdone":
		listTasksByStatus("not done")
	case "inprogress":
		listTasksByStatus("in progress")
	default:
		fmt.Println("Unknown command:", command)
	}
}
