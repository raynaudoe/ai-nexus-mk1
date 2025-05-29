mod stack;

use stack::Stack;

fn main() {
    let mut stack: Stack<i32> = Stack::new();

    stack.push(1);
    stack.push(2);
    stack.push(3);

    println!("Stack: {:?}", stack);
    println!("Is stack empty? {}", stack.is_empty());

    println!("Popped: {:?}", stack.pop());
    println!("Popped: {:?}", stack.pop());
    println!("Popped: {:?}", stack.pop());
    println!("Popped: {:?}", stack.pop());

    println!("Is stack empty? {}", stack.is_empty());
}
