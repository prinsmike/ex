package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("20060102+15-04-05", "20130217+19-36-47")
	if err != nil {
		panic(err)
	}
	fmt.Println("Year:", t.Year())
	fmt.Println("Month:", t.Month())
	fmt.Println("Day:", t.Day())
	fmt.Println("Hour:", t.Hour())
	fmt.Println("Minute:", t.Minute())
	fmt.Println("Seconds:", t.Second())
	fmt.Println("Weekday:", t.Weekday())
	fmt.Println("Local:", t.Local())
	fmt.Println("Location:", t.Location())
	fmt.Println("String:", t.String())
	fmt.Println("UTC:", t.UTC())
	fmt.Println("Unix:", t.Unix())
}
