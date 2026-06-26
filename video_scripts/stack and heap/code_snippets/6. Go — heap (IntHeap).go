// 6. Go — heap (IntHeap)
package main

import (
	"container/heap"
	"fmt"
)

// считаем, что IntHeap уже реализован и совместим с container/heap
// под капотом там хранится слайс чисел и поддерживается свойство кучи

func main() {
	h := &IntHeap{}
	heap.Init(h)

	heap.Push(h, 30)
	heap.Push(h, 10)
	heap.Push(h, 20)

	fmt.Println(*h) // [10 30 20]

	minValue := heap.Pop(h).(int)
	fmt.Println(minValue) // 10

	fmt.Println(*h) // [20 30]
}