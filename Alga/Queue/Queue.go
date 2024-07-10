package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		panic("Очередь пуста")
	}
	removed := q.items[0]
	q.items = q.items[1:]

	return removed
}

func (q *Queue) Peek() int {
	if len(q.items) == 0 {
		panic("Очередь пуста")
	}
	last := q.items[0]
	return last
}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	fmt.Println(queue.Peek())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println(queue.Dequeue())
}
