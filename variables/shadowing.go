// the b,a := 2,3 will create a new a variable that shadows the other a IF that line is in a nested block
package main

import "fmt"

func main() {
	var a int = 1

	// comment out the { and } to change the behavior of the program
	{
		b, a := 2, 3
		fmt.Printf("%d %d\n", a, b) // prints 3 2
	}
	fmt.Printf("%d\n", a) // print 1
}
