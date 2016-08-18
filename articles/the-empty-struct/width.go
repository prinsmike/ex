package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// We can discover the width of any value, and thus the width of its type
	// using the unsafe.Sizeof() function.

	var i32 int32
	var i64 int64
	var s string
	var c complex128
	fmt.Println(unsafe.Sizeof(i32)) // prints 4
	fmt.Println(unsafe.Sizeof(i64)) // prints 8
	fmt.Println(unsafe.Sizeof(s))   // prints 16
	fmt.Println(unsafe.Sizeof(c))   // prints 16

	// The width of an array type is a multiple of its element type.
	var a [3]uint32
	fmt.Println(unsafe.Sizeof(a)) // prints 12

	// Struct width is the sum of the with of its constituent types, plus padding.
	type S struct {
		a uint16
		b uint32
	}
	var s2 S
	fmt.Println(unsafe.Sizeof(s2)) // prints 8, not 6
}
