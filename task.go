package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

const taskFile = "tasks.json"

func readTasks() ([]Task, error) {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeTasks(tasks []Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}
