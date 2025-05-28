package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// Helper function to execute the todo CLI
func executeTodoCLI(args ...string) (string, error) {
	cmd := exec.Command("./todo", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func TestAddTodo(t *testing.T) {
	// Test adding a todo
	out, err := executeTodoCLI("add", "Buy groceries")
	if err != nil {
		t.Fatalf("Error adding todo: %v, output: %s", err, out)
	}

	if !strings.Contains(out, "Added todo: Buy groceries") {
		t.Errorf("Expected 'Added todo: Buy groceries', got: %s", out)
	}
}

func TestCompleteTodo(t *testing.T) {
	// First, add a todo
	out, err := executeTodoCLI("add", "Walk the dog")
	if err != nil {
		t.Fatalf("Error adding todo: %v, output: %s", err, out)
	}

	// Extract the ID of the added todo.  Assumes the ID is an integer.
	parts := strings.Split(out, " ")
	var id string
	for i, p := range parts {
		if p == "id:" {
			id = parts[i+1]
			break
		}
	}

	if id == "" {
		t.Fatalf("Could not find todo id in output: %s", out)
	}

	// Then, complete the todo
	out, err = executeTodoCLI("complete", id)
	if err != nil {
		t.Fatalf("Error completing todo: %v, output: %s", err, out)
	}

	if !strings.Contains(out, "Completed todo: "+id) {
		t.Errorf("Expected 'Completed todo: "+id+"', got: %s", out)
	}
}

func TestDeleteTodo(t *testing.T) {
	// First, add a todo
	out, err := executeTodoCLI("add", "Pay bills")
	if err != nil {
		t.Fatalf("Error adding todo: %v, output: %s", err, out)
	}

	// Extract the ID of the added todo
	parts := strings.Split(out, " ")
	var id string
	for i, p := range parts {
		if p == "id:" {
			id = parts[i+1]
			break
		}
	}

	if id == "" {
		t.Fatalf("Could not find todo id in output: %s", out)
	}

	// Then, delete the todo
	out, err = executeTodoCLI("delete", id)
	if err != nil {
		t.Fatalf("Error deleting todo: %v, output: %s", err, out)
	}

	if !strings.Contains(out, "Deleted todo: "+id) {
		t.Errorf("Expected 'Deleted todo: "+id+"', got: %s", out)
	}
}
