package main

import (
	"fmt"
	"strconv"
)

// Task struct
type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

// TaskList struct
type TaskList struct {
	tasks   map[int]Task
	nextID  int
}

// NewTaskList creates a new TaskList
func NewTaskList() *TaskList {
	return &TaskList{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

// AddTask adds a new task to the TaskList
func (tl *TaskList) AddTask(title, description string) int {
	task := Task{
		ID:          tl.nextID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	tl.tasks[tl.nextID] = task
	l.nextID++
	return task.ID
}

// DeleteTask deletes a task from the TaskList
func (tl *TaskList) DeleteTask(id int) error {
	if _, ok := tl.tasks[id]; !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}
	delete(tl.tasks, id)
	return nil
}

// CompleteTask marks a task as complete
func (tl *TaskList) CompleteTask(id int) error {
	task, ok := tl.tasks[id]
	if !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}
	task.Completed = true
	tl.tasks[id] = task
	return nil
}

// GetTask gets a task from the TaskList
func (tl *TaskList) GetTask(id int) (Task, error) {
	task, ok := tl.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("task with ID %d not found", id)
	}
	return task, nil
}

// GetAllTasks gets all tasks from the TaskList
func (tl *TaskList) GetAllTasks() []Task {
	tasks := make([]Task, 0, len(tl.tasks))
	for _, task := range tl.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// printTasks prints all tasks
func (tl *TaskList) printTasks() {
	for _, task := range tl.tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %t\n", task.ID, task.Title, task.Description, task.Completed)
	}
}

func main() {
	taskList := NewTaskList()

	for {
		var command string
		fmt.Print("Enter command (add, delete, complete, list, exit): ")
		fmt.Scanln(&command)

		switch command {
		case "add":
			var title, description string
			fmt.Print("Enter title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter description: ")
			fmt.Scanln(&description)
			id := taskList.AddTask(title, description)
			fmt.Printf("Task added with ID: %d\n", id)
		case "delete":
			var idStr string
			fmt.Print("Enter task ID to delete: ")
			fmt.Scanln(&idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			err = taskList.DeleteTask(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted")
			}
		case "complete":
			var idStr string
			fmt.Print("Enter task ID to complete: ")
			fmt.Scanln(&idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			err = taskList.CompleteTask(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task completed")
			}
		case "list":
			taskList.printTasks()
		case "exit":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
