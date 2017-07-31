package main

import "fmt"

func main() {

	m := make(map[int]string)
	m[3] = "Hello"
	m[7] = ", "
	m[22] = "world"
	m[30] = "!"

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	fmt.Printf("%+v\n", keys)
}
