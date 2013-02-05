// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"sync"
)

func main() {
	race()
	correct()
	alsoCorrect()
}

// This function has a data race and may print "55555", or something else.
func race() {
	fmt.Println("Data race:")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i) // The variable i is shared by six (6) goroutines.
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
}

func correct() {
	fmt.Println("Correct:")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) { // Use a local variable.
			fmt.Print(n)
			wg.Done()
		}(i)
	}
	wg.Wait() // Wait for all five goroutines to finish.
	fmt.Println()
}

func alsoCorrect() {
	fmt.Println("Also correct:")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		n := i // Create a unique variable for each closure.
		go func() {
			fmt.Print(n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
}
