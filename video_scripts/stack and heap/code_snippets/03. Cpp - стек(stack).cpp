// 3. Cpp — стек(stack)
#include <stack>
#include <iostream>

int main() {
    std::stack<int> stack;

    stack.push(10);
    stack.push(20);
    stack.push(30);

    int top = stack.top();
    stack.pop();
    std::cout << top << '\n'; // 30

    stack.push(40);

    int next_top = stack.top();
    stack.pop();
    std::cout << next_top << '\n'; // 40
}