package main

import (
	"fmt"

	"github.com/mailgun/holster"
)

func main() {
	queue := holster.NewPriorityQueue()

	queue.Push(&holster.PQItem{
		Value:    "thing3",
		Priority: 3,
	})
	queue.Push(&holster.PQItem{
		Value:    "thing1",
		Priority: 1,
	})
	queue.Push(&holster.PQItem{
		Value:    "thing2",
		Priority: 2,
	})

	var item *holster.PQItem

	item = queue.Pop()

	fmt.Printf("Item 1: %s\n", item.Value.(string))

	item = queue.Pop()
	fmt.Printf("Item 2: %s\n", item.Value.(string))
}
