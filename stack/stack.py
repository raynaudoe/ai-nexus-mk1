class Stack:
    def __init__(self):
        self._items = []

    def push(self, item):
        self._items.append(item)

    def pop(self):
        if not self.is_empty():
            return self._items.pop()
        else:
            return None  # Or raise an exception

    def is_empty(self):
        return len(self._items) == 0