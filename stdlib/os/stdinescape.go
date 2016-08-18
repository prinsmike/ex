package main

import (
	"bufio"
	"encoding/hex"
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
		line = EscapeCtrl(line)

		if string(line) == "exit\n" {
			fmt.Println("bye")
			return
		}

		fmt.Printf("Out %d: %s", i, string(line))
		i++
	}
}

func EscapeCtrl(ctrl []byte) (esc []byte) {
	u := []byte(`\u0000`)
	for i, ch := range ctrl {
		if ch <= 31 {
			if esc == nil {
				esc = append(make([]byte, 0, len(ctrl)+len(u)), ctrl[:i]...)
			}
			esc = append(esc, u...)
			hex.Encode(esc[len(esc)-2:], ctrl[i:i+1])
			continue
		}
		if esc != nil {
			esc = append(esc, ch)
		}
	}
	if esc == nil {
		return ctrl
	}
	return esc
}
