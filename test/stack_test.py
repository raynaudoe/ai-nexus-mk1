import unittest

class Stack:
    def __init__(self):
        self.items = []

    def is_empty(self):
        return len(self.items) == 0

    def push(self, item):
        self.items.append(item)

    def pop(self):
        if not self.is_empty():
            return self.items.pop()
        else:
            return None

class TestStack(unittest.TestCase):

    def setUp(self):
        self.stack = Stack()

    def test_is_empty(self):
        self.assertTrue(self.stack.is_empty())
        self.stack.push(1)
        self.assertFalse(self.stack.is_empty())

    def test_push(self):
        self.stack.push(10)
        self.stack.push(20)
        self.assertEqual(self.stack.items, [10, 20])

    def test_pop(self):
        self.stack.push(1)
        self.stack.push(2)
        self.assertEqual(self.stack.pop(), 2)
        self.assertEqual(self.stack.pop(), 1)
        self.assertTrue(self.stack.is_empty())
        self.assertIsNone(self.stack.pop())

    def test_multiple_pops_on_empty_stack(self):
        """Test that multiple pops on empty stack consistently return None"""
        self.assertIsNone(self.stack.pop())
        self.assertIsNone(self.stack.pop())
        self.assertIsNone(self.stack.pop())
        self.assertTrue(self.stack.is_empty())

    def test_mixed_operations(self):
        """Test complex sequences of push and pop operations"""
        # Test interleaved push/pop operations
        self.stack.push(1)
        self.stack.push(2)
        self.assertEqual(self.stack.pop(), 2)
        self.stack.push(3)
        self.stack.push(4)
        self.assertEqual(self.stack.pop(), 4)
        self.assertEqual(self.stack.pop(), 3)
        self.assertEqual(self.stack.pop(), 1)
        self.assertTrue(self.stack.is_empty())
        
        # Test push after emptying
        self.stack.push(5)
        self.assertFalse(self.stack.is_empty())
        self.assertEqual(self.stack.pop(), 5)

    def test_large_dataset(self):
        """Test stack behavior with many elements"""
        # Push many elements
        for i in range(100):
            self.stack.push(i)
        
        self.assertFalse(self.stack.is_empty())
        
        # Pop all elements and verify LIFO order
        for i in range(99, -1, -1):
            self.assertEqual(self.stack.pop(), i)
        
        self.assertTrue(self.stack.is_empty())
        self.assertIsNone(self.stack.pop())

    def test_duplicate_elements(self):
        """Test stack behavior with duplicate elements"""
        self.stack.push(5)
        self.stack.push(5)
        self.stack.push(5)
        
        self.assertEqual(self.stack.pop(), 5)
        self.assertEqual(self.stack.pop(), 5)
        self.assertEqual(self.stack.pop(), 5)
        self.assertTrue(self.stack.is_empty())

    def test_different_data_types(self):
        """Test stack with various data types"""
        elements = [1, "string", [1, 2, 3], {"key": "value"}, None]
        
        # Push different types
        for element in elements:
            self.stack.push(element)
        
        # Pop in reverse order
        for element in reversed(elements):
            self.assertEqual(self.stack.pop(), element)
        
        self.assertTrue(self.stack.is_empty())

    def test_edge_case_operations(self):
        """Test various edge cases"""
        # Test single element
        self.stack.push(42)
        self.assertFalse(self.stack.is_empty())
        self.assertEqual(self.stack.pop(), 42)
        self.assertTrue(self.stack.is_empty())
        
        # Test alternating single operations
        for i in range(5):
            self.stack.push(i)
            self.assertEqual(self.stack.pop(), i)
            self.assertTrue(self.stack.is_empty())

if __name__ == '__main__':
    unittest.main()
