package main

import (
	"fmt"
	"strconv"
)

func deleteTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: delete <task_id>")
		return
	}

	taskID := args[0]

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
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := writeTasks(tasks); err != nil {
				fmt.Println("Error writing tasks:", err)
				return
			}
			fmt.Println("Task deleted:", task.Name)
			return
		}
	}

	fmt.Println("Task not found:", taskID)
}
