// 8. rust — heap (BinaryHeap)
use std::cmp::Reverse;
use std::collections::BinaryHeap;

fn print_heap(heap: &BinaryHeap<Reverse<i32>>) {
    let mut copy = heap.clone();

    while let Some(Reverse(value)) = copy.pop() {
        print!("{} ", value);
    }
    println!();
}

fn example_heap() {
    let mut heap = BinaryHeap::new();

    heap.push(Reverse(30));
    heap.push(Reverse(10));
    heap.push(Reverse(20));

    print_heap(&heap); // 10 20 30 

    let top = heap.pop().unwrap();
    println!("{}", top.0); // 10

    print_heap(&heap); // 20 30 
}