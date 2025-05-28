ackage main

import (
	"fmt"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks []Task
var nextID = 1

func main() {
	for {
		fmt.Println("\nTodo List Application")
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Complete")
		fmt.Println("3. Delete Task")
		fmt.Println("4. List Tasks")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			markTaskComplete()
		case 3:
			deleteTask()
		case 4:
			listTasks()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addTask() {
	var description string
	fmt.Print("Enter task description: ")
	fmt.Scanln(&description)

	task := Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added successfully!")
}

func markTaskComplete() {
	var id int
	fmt.Print("Enter task ID to mark as complete: ")
	fmt.Scanln(&id)

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as complete!")
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask() {
	var id int
	fmt.Print("Enter task ID to delete: ")
	fmt.Scanln(&id)

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("Task not found.")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}

	fmt.Println("\nTasks:")
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("%d: %s (%s)\n", task.ID, task.Description, status)
	}
}
