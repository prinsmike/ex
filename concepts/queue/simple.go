// A simple queue implementation.

package main

import "fmt"

func main() {
	var v int
	q := NewQ()
	fmt.Println("Pushing...")
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Printf("Queue Length: %d\n", q.Len())
	fmt.Printf("Queue: %#v\n", q)
	v = q.Peek(0)
	fmt.Printf("Peeking: %d\n", v)
	v = q.Pop()
	fmt.Printf("Popped: %d\n", v)
	fmt.Println("Pushing...")
	q.Push(4)
	v = q.Peek(0)
	fmt.Printf("Peeking: %d\n", v)
	v = q.Pop()
	fmt.Printf("Popped: %d\n", v)
	fmt.Printf("Queue Length: %d\n", q.Len())
	fmt.Printf("Queue: %#v\n", q)
	v = q.Peek(0)
	fmt.Printf("Peeking: %d\n", v)
	v = q.Peek(q.Len() - 1)
	fmt.Printf("Peeking at last item: %d\n", v)
	fmt.Println("Clearing the queue...")
	q.Clear()
	fmt.Printf("Queue: %#v\n", q)
}

type Q []int

func NewQ() Q {
	return Q{}
}

func (q *Q) Push(val int) {
	*q = append(*q, val)
}

func (q *Q) Pop() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q Q) Peek(index int) int {
	return q[index]
}

func (q Q) Len() int {
	return len(q)
}

func (q Q) IsEmpty() bool {
	if len(q) == 0 {
		return true
	}
	return false
}

func (q *Q) Clear() {
	*q = []int{}
}
