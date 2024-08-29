package main

import "fmt"

func listTasks() {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d, Name: %s, Status: %s\n", task.ID, task.Name, task.Status)
	}
}

func listTasksByStatus(status string) {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	for _, task := range tasks {
		if task.Status == status {
			fmt.Printf("ID: %d, Name: %s, Status: %s\n", task.ID, task.Name, task.Status)
		}
	}
}
