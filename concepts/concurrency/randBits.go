// http://www.nada.kth.se/~snilsson/concurrency/
package main

import "fmt"

func main() {
	ch := RandomBits()
	for i := 10; i > 0; i-- {
		fmt.Print(<-ch)
	}
	fmt.Println()
}

// RandomBits returns a channel that produces a random sequence of bits.
func RandomBits() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0: // note: no statement
			case ch <- 1:
			}
		}
	}()
	return ch
}
