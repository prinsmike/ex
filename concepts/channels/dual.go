package main

import "fmt"

func main() {
	n := 3
	in := make(chan int)
	out := make(chan int)

	go square(in, out)

	in <- n
	fmt.Println(<-out)
}

func square(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	num := <-in

	result := num * num

	out <- result
}
