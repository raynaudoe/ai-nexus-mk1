use rustystack::Stack;
use rustystack::StackError;

fn main() {
    let mut stack = Stack::new(5);

    println!("Pushing elements onto the stack:");
    match stack.push(1) {
        Ok(_) => println!("Pushed 1"),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.push(2) {
        Ok(_) => println!("Pushed 2"),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.push(3) {
        Ok(_) => println!("Pushed 3"),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.push(4) {
        Ok(_) => println!("Pushed 4"),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.push(5) {
        Ok(_) => println!("Pushed 5"),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.push(6) {
        Ok(_) => println!("Pushed 6"),
        Err(e) => println!("Error: {:?}", e),
    }

    println!("\nPopping elements from the stack:");
    match stack.pop() {
        Ok(value) => println!("Popped {}", value),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.pop() {
        Ok(value) => println!("Popped {}", value),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.pop() {
        Ok(value) => println!("Popped {}", value),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.pop() {
        Ok(value) => println!("Popped {}", value),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.pop() {
        Ok(value) => println!("Popped {}", value),
        Err(e) => println!("Error: {:?}", e),
    }
    match stack.pop() {
        Ok(value) => println!("Popped {}", value),
        Err(e) => println!("Error: {:?}", e),
    }
}
