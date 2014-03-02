package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Clean("/var/www"))
	fmt.Println(filepath.Clean("/var/"))
	fmt.Println(filepath.Clean("~/.bashrc"))
	fmt.Println(filepath.Clean("/foo/bar/"))
}