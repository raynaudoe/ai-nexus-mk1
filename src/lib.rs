#[derive(Debug)]
pub enum StackError {
    Overflow,
    Underflow,
pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}

#[derive(Debug)]
pub struct Stack {
    stack: Vec<i32>,
    capacity: usize,
pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}

impl Stack {
    pub fn new(capacity: usize) -> Self {
        Stack {
            stack: Vec::with_capacity(capacity),
            capacity,
        pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}
    pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}

    pub fn push(&mut self, value: i32) -> Result<(), StackError> {
        if self.stack.len() == self.capacity {
            return Err(StackError::Overflow);
        pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}
        self.stack.push(value);
        Ok(())
    pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}

    pub fn pop(&mut self) -> Result<i32, StackError> {
        if self.stack.is_empty() {
            return Err(StackError::Underflow);
        pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}
        Ok(self.stack.pop().unwrap())
    pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}
pub fn is_empty(&self) -> bool {
        self.stack.is_empty()
    }
}

impl Default for Stack {
    fn default() -> Self {
        Self::new(100)
    }
}
