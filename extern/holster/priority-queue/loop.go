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

	qlen := queue.Len()

	fmt.Printf("Queue Length: %d\n", qlen)

	for i := 0; i < qlen; i++ {
		item = queue.Pop()
		fmt.Printf("Item %d: %s\n", i, item.Value.(string))
	}
}
