// 7. cpp — heap (priority_queue)
#include <iostream>
#include <queue>
#include <vector>
#include <functional>

int main() {
    std::priority_queue<int, std::vector<int>, std::greater<int>> heap;

    heap.push(30);
    heap.push(10);
    heap.push(20);

    for (auto copy = heap; !copy.empty(); copy.pop()) {
        std::cout << copy.top() << ' ';
    }
    std::cout << "\n"; // 10 20 30

    int top = heap.top();
    heap.pop();
    std::cout << top << "\n"; // 10

     for (auto copy = heap; !copy.empty(); copy.pop()) {
        std::cout << copy.top() << ' ';
    }
    std::cout << "\n"; // 20 30
}