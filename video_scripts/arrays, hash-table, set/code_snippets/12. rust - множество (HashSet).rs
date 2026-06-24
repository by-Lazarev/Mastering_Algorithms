// 12. rust - множество (HashSet)
use std::collections::HashSet;

fn main() {
    let mut visited: HashSet<String> = HashSet::new();

    visited.insert("Alice".to_string());
    visited.insert("Bob".to_string());
    visited.insert("Alice".to_string());

    println!("{}", visited.contains("Alice")); // true
    println!("{}", visited.contains("Carol")); // false
}