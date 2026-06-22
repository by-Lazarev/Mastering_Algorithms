// 3. cpp - массив(vector)
#include <vector>
#include <iostream>

int main() {
    std::vector<int> nums = {10, 20, 30, 40};

    for (int i = 0; i < nums.size(); i++) {
        std::cout << nums[i] << ' ';
    }
    std::cout << '\n';

    std::cout << nums[1] << '\n'; // 20

    nums.push_back(50);
    nums.insert(nums.begin() + 1, 15);

    for (int i = 0; i < nums.size(); i++) {
        std::cout << nums[i] << ' ';
    }
    std::cout << '\n'; // 10 15 20 30 40 50
}