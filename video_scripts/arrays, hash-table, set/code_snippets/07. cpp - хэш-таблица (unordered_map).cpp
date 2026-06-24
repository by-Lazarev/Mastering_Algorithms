//7. cpp - хэш-таблица (unordered_map)
#include <unordered_map>
#include <iostream>
#include <string>

int main() {
    std::unordered_map<std::string, int> ages = {
        {"Alice", 28},
        {"Bob", 32},
        {"Carol", 25}
    };

    std::cout << ages["Bob"] << '\n'; // 32

    ages["Dave"] = 30; // добавление новой пары

    if (ages.find("Alice")!=ages.end()) {
        std::cout << ages["Alice"] << '\n'; // 28
    }
}