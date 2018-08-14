package main

import "fmt"

func main() {
	out := make(chan int)
	in := make(chan int)

	go square(in, out)
	go square(in, out)
	go square(in, out)

	in <- 1
	in <- 2
	in <- 3

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func square(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	num := <-in
	result := num * num
	out <- result
}
