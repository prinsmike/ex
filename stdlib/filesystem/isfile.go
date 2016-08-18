package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter a file name\n")
	var s string
	fmt.Scanf("%s", &s)
	fi, err := os.Stat(s)
	if err != nil {
		fmt.Printf("%s does not exist!\n", s)
	}
	if fi.IsDir() {
		fmt.Printf("%s is a directory\n", s)
	}
	mode := fi.Mode()
	if mode&os.ModeSymlink == os.ModeSymlink {
		fmt.Printf("%s is a symbolic link\n", s)
	}
}
