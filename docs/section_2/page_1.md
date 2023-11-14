#two sum

=== "Python"
    
    ``` py
    class Solution:
        def twoSum(self, nums: List[int], target: int) -> List[int]:
            stack = {}
    
            for index, num in enumerate(nums):
                compliment_number = target - num
                if compliment_number in stack:
                    return stack[compliment_number], index
                stack[num] = index
    ```

=== "Rust"
    
    ``` rust
    use std::collections::{HashSet, HashMap};
    
    impl Solution {
        pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
            let mut stack = HashMap::new();
    
            for (index, number) in nums.iter().enumerate() {
                let compliment = target - number;
    
                if let Some(&i) = stack.get(&compliment) {
                    return vec![i as i32, index as i32];
                }
                stack.insert(*number, index);
            }
            unreachable!()
        }
    }
    ```
