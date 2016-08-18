package main

import (
	"encoding/hex"
	"fmt"
)

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

func main() {
	body := []byte("Hello,\t世界\n\x12")
	fmt.Println(body, string(body))
	body = EscapeCtrl(body)
	fmt.Println(body, string(body))
	body = []byte("Hello, 世界")
	fmt.Println(body, string(body))
	body = EscapeCtrl(body)
	fmt.Println(body, string(body))
}
