package main

import "fmt"

func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Task name is required.")
		return
	}

	name := args[0]

	tasks, err := readTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	newID := len(tasks) + 1

	task := Task{
		ID:     newID,
		Name:   name,
		Status: "not done",
	}

	tasks = append(tasks, task)

	if err := writeTasks(tasks); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Println("Task added:", task.Name)
}
