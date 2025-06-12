from stack.stack import Stack

if __name__ == "__main__":
    stack = Stack()
    print(f"Is stack empty? {stack.is_empty()}")  # True

    stack.push(1)
    stack.push(2)
    stack.push(3)

    print(f"Is stack empty? {stack.is_empty()}")  # False

    print(f"Popped: {stack.pop()}")  # 3
    print(f"Popped: {stack.pop()}")  # 2
    print(f"Popped: {stack.pop()}")  # 1

    print(f"Is stack empty? {stack.is_empty()}")  # True
    print(f"Popped: {stack.pop()}")  # None