package main

import (
	"bufio"
	"fmt"
	"os"
)
type Task struct{
	description string
	done bool
}

var ToDoList []Task

func main() {
	scanner:=bufio.NewScanner(os.Stdin)

	//creating an infinite loop
	for{
		fmt.Println("\nTo-Do-Go: A Simple To-Do List written in Go")
		fmt.Println("1. List Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Mark Task as Done")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		input:=scanner.Text()
		switch input{
		case "1":
			ListTasks()
		case "2":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := scanner.Text()
			AddTask(description)
			fmt.Println("Task added.")
		case "3":
			ListTasks()
			fmt.Print("Enter task index to remove: ")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			RemoveTask(index - 1)
			fmt.Println("Task removed.")
		case "4":
			ListTasks()
			fmt.Print("Enter task index to mark as done: ")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			MarkTaskDone(index - 1)
			fmt.Println("Task marked as done.")
		case "5":
			fmt.Println("See ya!~")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")

		}
	}
	
}

func ListTasks(){
	if len(ToDoList)==0{
		fmt.Println("No tasks to display.")
		return
	}
	for i, task:=range ToDoList{
		status:="Pending"
		if task.done{
			status="Done"
		}
		fmt.Printf("%d. %s - %s\n", i+1, task.description, status)
	}
}

func AddTask(description string){
	task:=Task{description, false}
	ToDoList=append(ToDoList, task)
}

func RemoveTask(index int) {
	if index >= 0 && index < len(ToDoList) {
		ToDoList = append(ToDoList[:index], ToDoList[index+1:]...)
	} else {
		fmt.Println("Invalid task index.")
	}
}

func MarkTaskDone(index int) {
	if index >= 0 && index < len(ToDoList) {
		ToDoList[index].done = true
	} else {
		fmt.Println("Invalid task index.")
	}
}