package main

import "fmt"

func main() {
	n := 3

	out := make(chan int)

	go square(n, out)

	fmt.Println(<-out)
}

func square(num int, out chan<- int) {
	result := num * num

	out <- result
}
