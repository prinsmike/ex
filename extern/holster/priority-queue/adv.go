package main

import (
	"fmt"

	"github.com/mailgun/holster"
)

type List []*holster.PQItem

func main() {
	queue := holster.NewPriorityQueue()
	list := List{
		&holster.PQItem{Value: "Item 5-1", Priority: 5},
		&holster.PQItem{Value: "Item 3", Priority: 3},
		&holster.PQItem{Value: "Item 2", Priority: 2},
		&holster.PQItem{Value: "Item 1", Priority: 1},
		&holster.PQItem{Value: "Item 5-2", Priority: 5},
		&holster.PQItem{Value: "Item 5-3", Priority: 5},
	}

	fmt.Println("Adding queue items...")
	for _, v := range list {
		fmt.Printf("Pushing item: %s\n", v.Value.(string))
		queue.Push(v)
	}

	fmt.Printf("Queue Length: %d\n", queue.Len())

	peek := queue.Peek()
	fmt.Printf("Peeking: %s\n", peek.Value.(string))

	fmt.Println("Adding a new item to the end of the queue...")
	queue.Push(&holster.PQItem{Value: "Last Item", Priority: queue.Len() + 1})
	fmt.Println("Adding a new item with priority 5...")
	queue.Push(&holster.PQItem{Value: "Item 5-4", Priority: 5})
	fmt.Printf("Queue Length: %d\n", queue.Len())

	var item *holster.PQItem
	qlen := queue.Len()
	for i := 0; i < qlen; i++ {
		item = queue.Pop()
		fmt.Printf("Popping item: %s\n", item.Value.(string))
	}
}
