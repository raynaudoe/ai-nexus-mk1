#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_push() {
        let mut stack = Stack::new();
        stack.push(1);
        assert_eq!(stack.size(), 1);
    }

    #[test]
    fn test_pop() {
        let mut stack = Stack::new();
        stack.push(1);
        let value = stack.pop();
        assert_eq!(value, Some(1));
        assert_eq!(stack.size(), 0);
    }

    #[test]
    fn test_is_empty() {
        let mut stack = Stack::new();
        assert_eq!(stack.is_empty(), true);
        stack.push(1);
        assert_eq!(stack.is_empty(), false);
    }

    #[test]
    fn test_pop_empty() {
        let mut stack = Stack::new();
        let value = stack.pop();
        assert_eq!(value, None);
    }
}
