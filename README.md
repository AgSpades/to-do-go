# To-Do List Application

A simple command-line To-Do list application written in Go. This app allows you to add, remove, mark tasks as done, and list all tasks interactively. It provides a simple CLI interface to manage your tasks.

## Features

- **Add a Task**: Add a new task to the to-do list.
- **Remove a Task**: Remove a task by specifying its index.
- **Mark a Task as Done**: Mark any task as completed.
- **List Tasks**: View all tasks with their current status (Pending or Done).
- **Exit**: Exit the application.

## Installation

To use the To-Do List Application, you need to have Go installed on your machine. If you don't have Go installed, you can follow the instructions on the official Go website: [https://golang.org/doc/install](https://golang.org/doc/install).

### Steps to Install:

1. Clone the repository:

   ```bash
   git clone https://github.com/agspades/to-do-go.git
   cd to-do-go
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

3. The app will prompt you with a menu. Choose an option and follow the instructions.

## Usage

After running the application, you will see a menu with the following options:

1. **List Tasks**: Display all the tasks in the to-do list.
2. **Add Task**: Add a new task by entering a description.
3. **Remove Task**: Remove a task by its index.
4. **Mark Task as Done**: Mark a task as completed by its index.
5. **Exit**: Exit the application.

### Example:

```bash
To-Do-Go: A Simple To-Do List written in Go
1. List Tasks
2. Add Task
3. Remove Task
4. Mark Task as Done
5. Exit
Choose an option: 2
Enter task description: Buy groceries
Task added.

To-Do-Go: A Simple To-Do List written in Go
1. List Tasks
2. Add Task
3. Remove Task
4. Mark Task as Done
5. Exit
Choose an option: 1
1. Buy groceries - Pending
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
