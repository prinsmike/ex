package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	h := sha1.New()
	io.WriteString(h, os.Args[1])
	fmt.Printf("%s\n", os.Args[1])
	fmt.Printf("%x\n", h.Sum(nil))
}
