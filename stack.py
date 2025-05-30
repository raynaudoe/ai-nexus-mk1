class Stack:
    def __init__(self):
        self._items = []

    def push(self, item):
        self._items.append(item)

    def pop(self):
        if not self.isEmpty():
            return self._items.pop()
        else:
            return None

    def isEmpty(self):
        return len(self._items) == 0

# Demonstration
if __name__ == '__main__':
    stack = Stack()
    print(f"Stack is empty: {stack.isEmpty()}")
    stack.push(1)
    stack.push(2)
    stack.push(3)
    print(f"Stack is empty: {stack.isEmpty()}")
    print(f"Popped: {stack.pop()}")
    print(f"Popped: {stack.pop()}")
    print(f"Popped: {stack.pop()}")
    print(f"Stack is empty: {stack.isEmpty()}")
    print(f"Popped from empty stack: {stack.pop()}")