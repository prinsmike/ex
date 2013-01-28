package main

import (
	"fmt"
)

func main() {
	t := []int{1, 2, 3, 4}
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", s)
}
