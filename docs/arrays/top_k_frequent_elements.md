# Название задания

[Ссылка с заданием на литкоде](https://leetcode.com/problems/top-k-frequent-elements/)


## Описание
Дан массив чисел `nums` и целое число `k`. Необходимо вернуть `k` самых часто встречающихся эллементов. Вернуть значения можно в любом порядке. Число `k` всегда меньше или равно количеству уникальных значений в `nums` 

**Пример 1:**  
> **Входные данные**: nums = [1, 1, 1, 2, 2, 3], k = 2  
> **Выходные данные**: [1, 2]  
> **Объяснение**: Числа по частоте вхождения: [1, 2, 3] - возвращаем первые 2 числа  

**Пример 2:**
> **Входные данные**: nums = [1], k = 1  
> **Выходные данные**: [1]  

**Пример 3:**
> **Входные данные**: nums = [1, 2, 3, 2], k = 1  
> **Выходные данные**: [2]  

---

## Идея решения
Увидев эту задачу в первый раз я подумал, что хитрое решение через приоритетные очереди будет эффективным. Оказалось, что решение через сортировку более эффективное, его и разберём, а второе решение я тоже добавлю ниже.

В целом оба подхода похожи. Сначала мы считаем количество вхождений каждого элемента. Задача была бы проще если бы нам нужно было вернуть только один, самый часто встречающийся элемент, но нам нужно `k` элементов, так что считаем все.

Посчитав вхождения, сортируем их по количеству вхождений и возвращаем первые `k` чисел.


## Алгоритм решения
В первой части нашего алгоритма мы итерируем массив и инкрементируем значение для встречающихся эллементов. Такое уже было в другой задаче этого же раздела.

Во второй части мы сортируем наши значения, при этом сортировка нужно по убыванию значений, а не ключей.
Далее возвращаем `k` первых ключей.

Таким образом наш алгоритм выглядит так:
1. создаём и заполняем хэш-таблицу ключи - уникальные числа, значения - кол-во их вхождений в массив
2. приводим нашу таблицу к списку кортежей(ключ, значение) и сортируем его по значениям
3. сохраняем `k` чисел(ключи хэш-таблицы) в результатирующий список
4. возвращаем результат в любом порядке

## Сложность

* По времени: $O(N + U log U)$ *N = len(nums), U = len(set(nums))*

Подсчёт элементов требует обратиться к каждому числу один раз, встравка в хэш-таблицу занимает константное время.

Преобразование хэш-таблицы и выборка `k` элементов занимает в худшем случае $O(U)$, где `U` - количество уникальных чисел в нашем массиве. Поскольку `U` всегда будет меньше чем `N` мы можем не учитывать этот этап, так как `N` мы уже учли в первой части. 

Значительное влияние на скорость оказывает сортировка, которая будет выполняться за $O(U log U)$. Можно было бы оптимизировать алгоритм выполняя сортировку в момент вставки в словарь, но размер решения и читаемость кода сильно измениться при таком подходе.

* По памяти: $O(U)$ *U = len(set(nums))*

При решении мы 4 раза выделяем память равную количеству уникальных элементов в массиве. 1 - для `counted_nums`, 2 - для `counted_vec`, 3 - для сортированного списка и 4 - для `result`. В упрощённом виде мы представляем это как $O(U)$.


## Код решения:

=== "Python"
    
    ``` py
    class Solution:
        def topKFrequent(self, nums: List[int], k: int) -> List[int]:
            counted_nums = {}

            for num in nums:
                counted_nums[num] = counted_nums.get(num, 0) + 1
            
            result = []
            counted_vec = counted_nums.items()
            for num, _ in sorted(counted_vec, key=lambda x: -x[1]):
                result.append(num)
                if len(result) >= k:
                    break

            return result
    ```

=== "Rust"
    
    ``` rust
    use std::collections::{HashMap, BinaryHeap};

    impl Solution {
        pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
            let mut counted_nums = HashMap::new();

            for num in nums {
                *counted_nums.entry(num).or_insert(1) += 1;
            }

            let mut result: Vec<i32> = Vec::with_capacity(k as usize);
            let mut counted_vec: Vec<(i32, i32)> = counted_nums.into_iter().collect();
            counted_vec.sort_by(|a, b| b.1.cmp(&a.1));
            for (num, _) in counted_vec.iter() {
                result.push(*num);
                if result.len() >= k as usize {
                    break;
                }
            }
            result
        }
    } 
    ```

=== "Python с очередью"
    
    ``` py
    class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counted_nums = {}

        for num in nums:
            counted_nums[num] = counted_nums.get(num, 0) + 1
        
        heap = []
        for num, count in counted_nums.items():
            heapq.heappush(heap, (-count, num))
        
        return [heapq.heappop(heap)[1] for _ in range(k)]
    ```

=== "Rust с очередью"
    
    ``` rust
    use std::collections::{HashMap, BinaryHeap};

    impl Solution {
        pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
            let mut counted_nums = HashMap::new();

            for num in nums {
                *counted_nums.entry(num).or_insert(0) += 1;
            }

            let mut heap = BinaryHeap::new();
            for (num, count) in counted_nums {
                heap.push((count, num));
            }

            let mut result = Vec::with_capacity(k as usize);
            for _ in 0..k {
                result.push(heap.pop().unwrap().1);
            }
            result
        }
    }    
    ```

=== "Python - решение в одну строку"
    ``` py
    class Solution:
        topKFrequent = lambda _, nums, k: [num for num, _ in Counter(nums).most_common(k)]
    ```
*В Python используется мин-куча, в связи с чем сохраняются данные с отрицательным знаком.* 


