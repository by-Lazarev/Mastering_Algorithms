// 11. cpp - множество (set)
#include <unordered_set>
#include <iostream>
#include <string>

int main() {
    std::unordered_set<std::string> visited;

    visited.insert("Alice");
    visited.insert("Bob");
    visited.insert("Alice");

    std::cout << visited.contains("Alice") << '\n'; // 1
    std::cout << visited.contains("Carol") << '\n'; // 0
}