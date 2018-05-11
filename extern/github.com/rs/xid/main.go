package main

import (
	"fmt"

	"github.com/rs/xid"
)

func main() {
	guid := xid.New()
	fmt.Printf("GUID: %s\n", guid.String())
	fmt.Printf("Machine ID: %x\n", guid.Machine())
	fmt.Printf("Process ID: %d\n", guid.Pid())
	fmt.Printf("Time: %v\n", guid.Time())
	fmt.Printf("Counter: %d\n", guid.Counter())

}
