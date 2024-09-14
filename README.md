# Task Tracker CLI

A command-line interface (CLI) application to manage tasks efficiently. This project enables users to add, remove, update, and display tasks, storing them in a JSON file. It follows the principles outlined in the [Roadmap Project](https://roadmap.sh/projects/task-tracker).

## Features

- **Add Todo**: Create a new task with a description and status.
- **Remove Todo**: Delete a task by its index.
- **Display Todos**: Show all tasks with their details.
- **Update Todo Status**: Change the status of an existing task.

## Requirements

- Go 1.16 or later
- Compatible with major operating systems (e.g., Linux, macOS, Windows)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/bot69dude/Task-Tracker.git
   cd Task-Tracker
   ```

2. Build the application:
   ```bash
   go build -o task-tracker
   ```

## Usage

1. Run the application:
   ```bash
   ./task-tracker
   ```

2. The following options will be available:
   * **1**: Add Todo
   * **2**: Remove Todo
   * **3**: Display All Todos
   * **4**: Update Todo Status
   * **5**: Exit

3. Follow the prompts to manage your tasks.

## Data Storage

Tasks are saved in a `data.json` file located in the same directory as the executable. This file is automatically created and updated by the application.

## Project Structure

* `main.go`: The main application file containing the implementation of task management functions.
* `data.json`: Stores tasks in JSON format.

## Development

To contribute or modify the code:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-branch
   ```
3. Make your changes and commit:
   ```bash
   git add .
   git commit -m "Describe your changes"
   ```
4. Push to your forked repository:
   ```bash
   git push origin feature-branch
   ```
5. Open a pull request to merge changes.

## Testing

Run the tests using the following command:
```bash
go test ./...
```

## License

This project is open source and available under the [MIT License](LICENSE).