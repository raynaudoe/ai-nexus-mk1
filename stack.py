class Stack:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        if not self.is_empty():
            return self.items.pop()
        else:
            raise IndexError("Cannot pop from an empty stack")

    def is_empty(self):
        return len(self.items) == 0

# Example Usage:
if __name__ == "__main__":
    stack = Stack()
    stack.push(1)
    stack.push(2)
    stack.push(3)

    print("Stack is empty:", stack.is_empty())

    print("Popped item:", stack.pop())
    print("Popped item:", stack.pop())
    print("Popped item:", stack.pop())

    print("Stack is empty:", stack.is_empty())

    try:
        stack.pop()
    except IndexError as e:
        print(e)
