#[derive(Debug)]
pub enum StackError {
    Overflow,
    Underflow,
}

#[derive(Debug)]
pub struct Stack {
    stack: Vec<i32>,
    capacity: usize,
}

impl Stack {
    pub fn new(capacity: usize) -> Self {
        Stack {
            stack: Vec::with_capacity(capacity),
            capacity,
        }
    }

    pub fn push(&mut self, value: i32) -> Result<(), StackError> {
        if self.stack.len() >= self.capacity {
            return Err(StackError::Overflow);
        }
        self.stack.push(value);
        Ok(())
    }

    pub fn pop(&mut self) -> Result<i32, StackError> {
        if self.stack.is_empty() {
            return Err(StackError::Underflow);
        }
        Ok(self.stack.pop().unwrap())
    }
}
