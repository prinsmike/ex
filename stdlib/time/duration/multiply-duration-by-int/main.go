package main

import (
	"fmt"
	"time"
)

func main() {
	var seconds int64 = 300 // 5 minutes
	fmt.Printf("Three times five minutes: %s\n",
		(time.Duration(seconds)*time.Second)*time.Duration(3))

	fmt.Printf("Five times five minutes: %s\n",
		(time.Duration(seconds)*time.Second)*time.Duration(5))

	fmt.Printf("Twelve times five minutes: %s\n",
		(time.Duration(seconds)*time.Second)*time.Duration(12))

	fmt.Printf("Eighteen times five minutes: %s\n",
		(time.Duration(seconds)*time.Second)*time.Duration(18))
}
