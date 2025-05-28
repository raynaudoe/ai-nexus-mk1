ackage main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TaskManager manages tasks
type TaskManager struct {
	tasks   []Task
	storage *DataStorage
}

// NewTaskManager creates a new TaskManager
func NewTaskManager(storage *DataStorage) *TaskManager {
	return &TaskManager{
		tasks:   []Task{},
		storage: storage,
	}
}

// AddTask adds a new task
func (tm *TaskManager) AddTask(description string) {
	newTask := Task{
		ID:          len(tm.tasks) + 1,
		Description: description,
	}
	tm.tasks = append(tm.tasks, newTask)
	tm.storage.SaveTasks(tm.tasks)
}

// DeleteTask deletes a task
func (tm *TaskManager) DeleteTask(id int) error {
	index := id - 1
	if index < 0 || index >= len(tm.tasks) {
		return fmt.Errorf("invalid task ID")
	}

	tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)

	// Adjust IDs of remaining tasks
	for i := index; i < len(tm.tasks); i++ {
		tm.tasks[i].ID = i + 1
	}

	tm.storage.SaveTasks(tm.tasks)
	return nil
}

// ListTasks lists all tasks
func (tm *TaskManager) ListTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}
	for _, task := range tm.tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Description)
	}
}

// DataStorage handles data storage
type DataStorage struct {
	filePath string
}

// NewDataStorage creates a new DataStorage
func NewDataStorage(filePath string) *DataStorage {
	return &DataStorage{
		filePath: filePath,
	}
}

// LoadTasks loads tasks from file
func (ds *DataStorage) LoadTasks() ([]Task, error) {
	file, err := os.Open(ds.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, return an empty list
			return []Task{}, nil
		}
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	tasks := []Task{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Skipping invalid line: ", line)
			continue
		}
		id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Skipping invalid line: ", line)
			continue
		}
		description := strings.TrimSpace(parts[1])
		tasks = append(tasks, Task{ID: id, Description: description})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return tasks, nil
}

// SaveTasks saves tasks to file
func (ds *DataStorage) SaveTasks(tasks []Task) error {
	file, err := os.Create(ds.filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		_, err := fmt.Fprintf(writer, "%d: %s\n", task.ID, task.Description)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	return nil
}
