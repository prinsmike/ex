// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"time"
)

func main() {
	wait := Publish("Channels let goroutines communicate.", 5*time.Second)
	fmt.Println("Waiting for the news...")
	<-wait
	fmt.Println("The news is out, time to leave.")
}

// Publish prints text to stdout after the given time has expired.
// It closes the wait channel when the text has been published.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch) // broadcast - a closed channel sends a zero value forever
	}()
	return ch
}
