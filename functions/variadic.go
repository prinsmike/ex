package main

import (
	"fmt"
)

func variadicFunc(params ...string) {
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {

	arg1 := "Line 1"
	arg2 := "Line 2"
	arg3 := "Line 3"
	arg4 := "Line 4"

	argSlice := []string{"Line 5", "Line 6", "Line 7"}

	variadicFunc(arg1, arg2, arg3)
	variadicFunc(arg4)
	variadicFunc(argSlice...)
}
