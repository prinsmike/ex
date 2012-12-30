// Check if a value is in a slice.

package main

import (
	"fmt"
)

func InString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func InByte(slice []byte, value byte) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	s := []string{"Hello", "World", "Gopher"}
	b := []byte("abcdefghijklmnopqrstuvwxyz")
	fmt.Println(InString(s, "Gopher"))
	fmt.Println(InString(s, "Not in there"))
	fmt.Println(InByte(b, 'm'))
	fmt.Println(InByte(b, '5'))
}
