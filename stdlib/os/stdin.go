package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	i := 0
	for {
		fmt.Printf("In %d: ", i)
		line, err := in.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			return
		}

		if string(line) == "exit\n" {
			fmt.Println("bye")
			return
		}

		fmt.Printf("Out %d: %s", i, string(line))
		i++
	}
}
