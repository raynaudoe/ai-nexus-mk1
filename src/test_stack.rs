#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_push() {
        let mut stack = Stack::new();
        stack.push(1);
        assert_eq!(stack.items.len(), 1);
    }

    #[test]
    fn test_pop() {
        let mut stack = Stack::new();
        stack.push(1);
        let value = stack.pop();
        assert_eq!(value, Some(1));
        assert_eq!(stack.items.len(), 0);
    }

    #[test]
    fn test_is_empty() {
        let mut stack = Stack::new();
        assert_eq!(stack.is_empty(), true);
        stack.push(1);
        assert_eq!(stack.is_empty(), false);
    }
}
