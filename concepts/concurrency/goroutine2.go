// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"time"
)

// This program will, most likely, print both "Hello from main goroutine" and
// "Hello from another goroutine". They might be printed in any order.
func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	time.Sleep(time.Second) // wait 1 sec for other goroutine to finish

	// At this point the program execution stops and all
	// active goroutines are killed.
}
