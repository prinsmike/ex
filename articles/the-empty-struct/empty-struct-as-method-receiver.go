// Empty structs behave just like any other type.
// We can therefore use them as method receivers.

package main

import "fmt"

type S struct{}

func (s *S) addr() { fmt.Printf("%p\n", s) }

func main() {
	var a, b S
	a.addr() // 0x1beeb0 (exact address depends on Go version)
	b.addr() // 0x1beeb0 (exact address depends on Go version)
}
