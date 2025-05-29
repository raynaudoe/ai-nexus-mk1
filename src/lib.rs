/// A simple stack data structure.
///
/// # Examples
///
/// ```
/// let mut stack = Stack::new();
/// stack.push(1);
/// stack.push(2);
/// assert_eq!(stack.pop(), Some(2));
/// assert_eq!(stack.pop(), Some(1));
/// assert_eq!(stack.pop(), None);
/// ```
pub struct Stack<T> {
    data: Vec<T>,
}

impl<T> Stack<T> {
    /// Creates a new empty stack.
    pub fn new() -> Self {
        Stack {
            data: Vec::new(),
        }
    }

    /// Pushes an element onto the top of the stack.
    ///
    /// # Arguments
    ///
    /// * `item` - The element to push onto the stack.
    pub fn push(&mut self, item: T) {
        self.data.push(item);
    }

    /// Removes and returns the element at the top of the stack, if any.
    ///
    /// # Returns
    ///
    /// Removes and returns the element at the top of the stack, if any.
    ///
    /// # Returns
    ///
    /// An `Option` containing the element at the top of the stack, or `None` if the stack is empty.
    pub fn pop(&mut self) -> Option<T> {
        self.data.pop()
    }

    /// Returns a reference to the top element of the stack, without removing it.
    ///
    /// # Returns
    ///
    /// An `Option` containing a reference to the top element of the stack, or `None` if the stack is empty.
    pub fn peek(&self) -> Option<&T> {
        self.data.last()
    }

    /// Returns the number of elements in the stack.
    pub fn len(&self) -> usize {
        self.data.len()
    }

    /// Returns `true` if the stack is empty.
    pub fn is_empty(&self) -> bool {
        self.data.is_empty()
    }
}

#[cfg(test)]
mod tests {
    use super::Stack;

    #[test]
    fn push_pop_works() {
        let mut stack: Stack<i32> = Stack::new();
        stack.push(1);
        stack.push(2);
        assert_eq!(stack.pop(), Some(2));
        assert_eq!(stack.pop(), Some(1));
        assert_eq!(stack.pop(), None);
    }

    #[test]
    fn peek_works() {
        let mut stack: Stack<i32> = Stack::new();
        stack.push(1);
        assert_eq!(stack.peek(), Some(&1));
        stack.pop();
        assert_eq!(stack.peek(), None);
    }

    #[test]
    fn len_works() {
        let mut stack: Stack<i32> = Stack::new();
        assert_eq!(stack.len(), 0);
        stack.push(1);
        assert_eq!(stack.len(), 1);
        stack.push(2);
        assert_eq!(stack.len(), 2);
        stack.pop();
        assert_eq!(stack.len(), 1);
    }

    #[test]
    fn is_empty_works() {
        let mut stack: Stack<i32> = Stack::new();
        assert_eq!(stack.is_empty(), true);
        stack.push(1);
        assert_eq!(stack.is_empty(), false);
        stack.pop();
        assert_eq!(stack.is_empty(), true);
    }
}
