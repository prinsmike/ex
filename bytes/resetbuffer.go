package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	b1 := []byte("abcdef")
	b2 := []byte("ghijkl")

	for _, c := range b1 {
		buffer.WriteByte(c)
	}
	s1 := buffer.String()
	buffer.Reset()
	fmt.Println(s1)

	for _, c := range b2 {
		buffer.WriteByte(c)
	}
	s2 := buffer.String()
	fmt.Println(s2)
}
