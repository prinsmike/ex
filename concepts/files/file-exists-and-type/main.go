package main

import (
	"log"
	"os"
	"path"
)

func main() {
	curdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fp := path.Join(curdir, "test")
	if fi, err := os.Stat(fp); err == nil && !fi.IsDir() {
		log.Println("File exists and is not a directory.")
	} else {
		log.Println("File does not exist or is a directory.")
	}
}
