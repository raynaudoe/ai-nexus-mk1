import tkinter as tk
from task_manager import TaskManager

class TodoApp:
    def __init__(self, root):
        self.root = root
        self.root.title("Todo List App")
        self.task_manager = TaskManager()
        self.task_manager.load_tasks()
        self.setup_ui()

    def setup_ui(self):
        # TODO: Implement UI elements here
        pass

if __name__ == "__main__":
    root = tk.Tk()
    app = TodoApp(root)
    root.mainloop()