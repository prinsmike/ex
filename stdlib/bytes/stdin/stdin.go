package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(line)
		fmt.Printf("%s\n", string(line))
	}
}
