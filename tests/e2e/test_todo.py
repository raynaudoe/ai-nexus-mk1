import unittest
import subprocess
import json
import os

class TestTodoApp(unittest.TestCase):

    def setUp(self):
        # Define the command to run the todo app
        self.command = ["python", "main.py"]
        # Ensure tasks.json exists and is empty before each test
        with open("tasks.json", "w") as f:
            json.dump([], f)

    def tearDown(self):
        # Clean up any changes to tasks.json after each test
        with open("tasks.json", "w") as f:
            json.dump([], f)

    def run_command(self, input_str=""):
        process = subprocess.Popen(
            self.command,
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True
        )
        stdout, stderr = process.communicate(input=input_str)
        return stdout, stderr, process.returncode

    def test_add_task(self):
        # Test adding a task
        stdout, stderr, returncode = self.run_command("add Buy groceries\nlist\n")
        self.assertEqual(returncode, 0)
        self.assertIn("Task added", stdout)
        # Verify that the task is added to tasks.json
        with open("tasks.json", "r") as f:
            tasks = json.load(f)
        self.assertEqual(len(tasks), 1)
        self.assertEqual(tasks[0]["description"], "Buy groceries")

    def test_list_tasks(self):
        # Add a task first
        self.run_command("add Buy groceries\nlist\n")
        # Test listing tasks
        stdout, stderr, returncode = self.run_command("list\n")
        self.assertEqual(returncode, 0)
        self.assertIn("Buy groceries", stdout)

    def test_delete_task(self):
        # Add a task first
        self.run_command("add Buy groceries\nlist\n")
        # Delete the task
        stdout, stderr, returncode = self.run_command("delete 0\nlist\n")
        self.assertEqual(returncode, 0)
        self.assertIn("Task deleted", stdout)
        # Verify that tasks.json is empty
        with open("tasks.json", "r") as f:
            tasks = json.load(f)
        self.assertEqual(len(tasks), 0)

    def test_edit_task(self):
        # Add a task first
        self.run_command("add Buy groceries\nlist\n")
        # Edit the task
        stdout, stderr, returncode = self.run_command("edit 0 Read a book\nlist\n")
        self.assertEqual(returncode, 0)
        self.assertIn("Task edited", stdout)
        # Verify that the task is updated in tasks.json
        with open("tasks.json", "r") as f:
            tasks = json.load(f)
        self.assertEqual(tasks[0]["description"], "Read a book")

if __name__ == '__main__':
    unittest.main()
