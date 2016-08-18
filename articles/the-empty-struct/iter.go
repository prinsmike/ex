// Brad Fitzpatrick's iter package uses empty structs
// to create a slice of n 0-sized elements, suitable for
// ranging over.

package main

import (
	"fmt"
	"unsafe"

	"github.com/bradfitz/iter"
)

func main() {

	fmt.Println(unsafe.Sizeof(iter.N(10))) // prints 12 on the playground.

	// Print 0 - 9, inclusive, without causing any allocations.
	for i := range iter.N(10) {
		fmt.Println(i)
	}
}
