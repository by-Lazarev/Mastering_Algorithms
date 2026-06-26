// 2. Go — стек

package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	n := len(s.items) - 1
	value := s.items[n]
	s.items = s.items[:n]
	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) String() string {
	return fmt.Sprint(s.items)
}

func main() {
	var stack Stack

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	top, _ := stack.Pop()
	fmt.Println(top) // 30

	stack.Push(40)

	nextTop, _ := stack.Pop()
	fmt.Println(nextTop) // 40
}