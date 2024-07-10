package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Стек пуст")
	}
	lastIndex := len(s.items) - 1
	removed := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return removed
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		panic("Стек пуст")
	}
	lastIndex := len(s.items) - 1
	last := s.items[lastIndex]

	return last
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Peek())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.Pop())
}
