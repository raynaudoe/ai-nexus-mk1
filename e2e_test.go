package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// TestTodoAppE2E tests the end-to-end functionality of the todo list application.
func TestTodoAppE2E(t *testing.T) {
	// Build the application.
	cmdBuild := exec.Command("go", "build", ".")
	errBuild := cmdBuild.Run()
	if errBuild != nil {
		t.Fatalf("Failed to build the application: %v", errBuild)
	}

	// Define test cases.
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "AddTask",
			input:          "add Buy groceries\nMilk, eggs, bread",
			expectedOutput: "Task added",
		},
		{
			name:           "ViewTasks",
			input:          "view",
			expectedOutput: "Buy groceries",
		},
		{
			name:           "CompleteTask",
			input:          "complete 1",
			expectedOutput: "Task marked as complete",
		},
		{
			name:           "DeleteTask",
			input:          "delete 1",
			expectedOutput: "Task deleted",
		},
	}

	// Execute test cases.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Execute the command.
			cmdRun := exec.Command("./todo-list-app-2") // Assuming the executable is named after the directory
			cmdRun.Stdin = strings.NewReader(test.input)

			outputBytes, errRun := cmdRun.CombinedOutput()
			if errRun != nil {
				t.Fatalf("Failed to run the application: %v", errRun)
			}
			output := string(outputBytes)

			// Check the output.
			if !strings.Contains(output, test.expectedOutput) {
				t.Errorf("Test '%s' failed: expected to contain '%s', got '%s'", test.name, test.expectedOutput, output)
			}
		})
	}

	// Cleanup: Remove the executable.
	errRemove := exec.Command("rm", "todo-list-app-2").Run()
	if errRemove != nil {
		t.Fatalf("Failed to remove the executable: %v", errRemove)
	}
}
