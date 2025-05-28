ackage main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var nextID = 1

const filename = "tasks.json"

func main() {
	loadTasks()

	addTask("Buy groceries")
	addTask("Do laundry")
	completeTask(1)

	listTasks()

	saveTasks()
}

func loadTasks() {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
}

func saveTasks() {
	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func addTask(name string) {
	task := Task{ID: nextID, Name: name, Completed: false}
	tasks = append(tasks, task)
	nextID++
	fmt.Printf("Added task: %s\n", name)
}

func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Printf("Completed task: %s\n", tasks[i].Name)
			return
		}
	}
	fmt.Println("Task not found")
}

func listTasks() {
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("%d: %s (%s)\n", task.ID, task.Name, status)
	}
}
