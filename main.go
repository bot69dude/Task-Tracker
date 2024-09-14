package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var todos = make(map[int]Todo)
var idx int
var fileName = "data.json"

func (t *Todo) SetTodo(description string, status string) {
	now := time.Now()
	t.Description = description
	t.Status = status
	t.CreatedAt = now
	t.UpdatedAt = now
	todos[idx] = *t
	idx++

	jsonData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data successfully written to", fileName)
}

func LoadTodosFromFile() {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(fileData, &todos)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for k := range todos {
		if k >= idx {
			idx = k + 1
		}
	}

	fmt.Println("Todos loaded successfully from", fileName)
}

func RemoveTodo(index int) {
	LoadTodosFromFile()

	if _, exists := todos[index]; exists {
		delete(todos, index)

		newTodos := make(map[int]Todo)
		newIdx := 0
		for _, todo := range todos {
			newTodos[newIdx] = todo
			newIdx++
		}
		todos = newTodos
		idx = newIdx

		jsonData, err := json.MarshalIndent(todos, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}

		err = os.WriteFile(fileName, jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("JSON data successfully updated to", fileName)
	} else {
		fmt.Println("Todo not found.")
	}
}

func UpdateTodoStatus(index int, newStatus string) {
	LoadTodosFromFile()

	if todo, exists := todos[index]; exists {
		todo.Status = newStatus
		todo.UpdatedAt = time.Now()
		todos[index] = todo

		jsonData, err := json.MarshalIndent(todos, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}

		err = os.WriteFile(fileName, jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Todo status updated successfully.")
	} else {
		fmt.Println("Todo not found.")
	}
}

func DisplayAll() {
	if len(todos) > 0 {
		fmt.Println("List of all todos:")
		for i, todo := range todos {
			fmt.Printf("%d. Description: %s\nStatus: %s\nCreated At: %s\nLast Updated At: %s\n\n", i, todo.Description, todo.Status, todo.CreatedAt, todo.UpdatedAt)
		}
	} else {
		fmt.Println("No todos to display.")
	}
}

func main() {
	LoadTodosFromFile()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. For Adding Todo")
		fmt.Println("2. For Removing Todo")
		fmt.Println("3. For Displaying All Todos")
		fmt.Println("4. For Updating Todo Status")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		var choice int
		fmt.Sscanf(scanner.Text(), "%d", &choice)

		switch choice {
		case 1:
			fmt.Print("Please Enter Todo Description: ")
			scanner.Scan()
			description := scanner.Text()

			fmt.Print("Please Enter Status of TODO List: ")
			scanner.Scan()
			status := scanner.Text()

			var t Todo
			t.SetTodo(description, status)
			fmt.Println("Todo added successfully!")

		case 2:
			DisplayAll()
			fmt.Print("Enter the index of the todo to remove: ")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			RemoveTodo(index)

		case 3:
			DisplayAll()

		case 4:
			DisplayAll()
			fmt.Print("Enter the index of the todo to update: ")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			fmt.Print("Enter the new status: ")
			scanner.Scan()
			newStatus := scanner.Text()
			UpdateTodoStatus(index, newStatus)

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
