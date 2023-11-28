# Two sum

[Ссылка с заданием на литкоде](https://leetcode.com/problems/two-sum)


## Описание

Для заданого массива целых чисел `nums` и целого числа `target` необходимо вернуть два индекса в `nums`. Сумма чисел по этим индексам должна образовывать число `target`.

Гарантируется, что при любых тестовых данных есть только одно решение и оба числа не могут быть одинаковыми.

Ответ можно вернуть в любом порядке.

**Пример 1:**  
> **Входные данные**: nums = [2, 7, 11, 15], target = 9  
> **Выходные данные**:  [0, 1]  
> **Объяснение**: Поскольку nums[0] + nums[7] == 9, мы вернули [0, 1]

**Пример 2:**
> **Входные данные**: nums = [3, 2, 4], target = 6  
> **Выходные данные**:  [1, 2]

**Пример 3:**
> **Входные данные**: nums = [], target =    
> **Выходные данные**: []  

---

## Идея решения
Решение в лоб для данного примера будет представлять проверку всех комбинаций. Берём первое число в массиве и ищем, есть ли в массиве число равное разности `target` и текущему числу. В данном случае решение не будет эффективным, однако его можно взять за основу и ускорить используя хэш-таблицу.

## Алгоритм решения


## Сложность

* По времени

* По памяти

## Код решения:

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
