//8. rust - хэш-таблица (HashMap)

use std::collections::HashMap;

fn main() {
    let mut ages = HashMap::new();

    ages.insert(String::from("Alice"), 28);
    ages.insert(String::from("Bob"), 32);
    ages.insert(String::from("Carol"), 25);

    if let Some(age) = ages.get("Bob") {
        println!("{}", age); // 32
    }

    ages.insert(String::from("Dave"), 30);

    if ages.contains_key("Alice") {
        println!("{}", ages["Alice"]); // 28
    }
}