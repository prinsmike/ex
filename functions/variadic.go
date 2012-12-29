package main

import (
	"fmt"
)

func variadicFunc(params ...string) {
	for _, param := range params {
		fmt.Println(param)
	}
}

func variadicFunc2(param1 string, params ...string) {
	for _, param := range params {
		fmt.Println(param1, param)
	}
}

func main() {

	arg1 := "Line 1"
	arg2 := "Line 2"
	arg3 := "Line 3"
	arg4 := "Line 4"

	argSlice := []string{"Line 5", "Line 6", "Line 7"}

	argSlice2 := []string{"Line 8", "Line 9", "Line 10"}

	variadicFunc(arg1, arg2, arg3)
	variadicFunc(arg4)
	variadicFunc(argSlice...)
	variadicFunc2("This is", argSlice2...)
}
