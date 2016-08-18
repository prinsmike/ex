package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// The empty struct has a width of zero.
	var s struct{}

	fmt.Println(unsafe.Sizeof(s)) // prints 0

	// A struct comprised of empty structs also consumes no storage.
	type S struct {
		A struct{}
		B struct{}
	}

	var s2 S
	fmt.Println(unsafe.Sizeof(s2)) // prints 0

	// An array of empty structs consumes no storage.
	var x [1000000000]struct{}
	fmt.Println(unsafe.Sizeof(x)) // prints 0

	// Slices of empty structs consume only the space for their header.
	var x2 = make([]struct{}, 1000000000)
	fmt.Println(unsafe.Sizeof(x2)) // prints 12 in the playground.

	// The normal subslice, len, and cap builtins work as expected.
	var x3 = make([]struct{}, 100)
	var y = x3[:50]
	fmt.Println(len(y), cap(y)) // prints 50 100

	// You can take the address of a struct{} value when it is addressable.
	var a struct{}
	var b = &a
	fmt.Println(b) // prints &{}

	// The address of two struct{} values may be the same.
	var a2, b2 struct{}
	fmt.Println(&a2 == &b2) // true

	a3 := make([]struct{}, 10)
	b3 := make([]struct{}, 20)
	fmt.Println(&a3 == &b3)       // false, a3 and b3 are different slices.
	fmt.Println(&a3[0] == &b3[0]) // true, their backing arrays are the same.

	// Empty structs contain no fields, so can hold no data.
	// If empty structs hold no data, it is not possible to determine if two
	// struct{} values are different.
	a4 := struct{}{} // not the zero value, a real new struct{} instance.
	b4 := struct{}{}
	fmt.Println(a4 == b4) // true
}
