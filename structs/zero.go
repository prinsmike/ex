package main

import (
	"fmt"
)

type Zero struct {
	Text  string
	Int   int
	Float float64
	Bool  bool
	Slice []string
}

func main() {
	zero := &Zero{}
	fmt.Printf("%#v\n", zero)
}
