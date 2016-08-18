package main

import (
"fmt"
"strconv"
"os"
)

func main() {
	
	n, e := strconv.Atoi(os.Args[1])
	if e != nil {
		panic(e)
	}

	for i:=1; i<=20; i++ {
		if n == i {
			fmt.Printf("\x1b[33;1m%d * %d = %d (squared)\x1b[0m\n", n, i, n*i)
		} else {
			fmt.Printf("%d * %d = %d\n", n, i, n*i)
		}
	}	
}