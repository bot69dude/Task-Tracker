package main

import (
	"fmt"
	"strconv"
)

func updateTask(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: update <task_id> <new_status>")
		return
	}

	taskID := args[0]
	newStatus := args[1]

	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	id, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("Invalid task ID:", taskID)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = newStatus
			if err := writeTasks(tasks); err != nil {
				fmt.Println("Error writing tasks:", err)
				return
			}
			fmt.Println("Task updated:", task.Name)
			return
		}
	}

	fmt.Println("Task not found:", taskID)
}
