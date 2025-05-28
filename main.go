package main

import (
	"fmt"
	os"
)

type Task struct {
	ID          int
	Description string
}

const filename = "tasks.txt"

func main() {
	fmt.Println("Todo app in Go")
	// TODO: Implement CLI argument parsing and task management logic
}
