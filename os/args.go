package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough arguments.")
	}
	println("Hello", os.Args[1])
}
