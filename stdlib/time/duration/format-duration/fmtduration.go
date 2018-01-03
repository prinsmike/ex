package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Round(0).Add(-(3600 + 60 + 45) * time.Second)
	since := time.Since(t)
	fmt.Println(since)
	d := fmtDuration(since)
	fmt.Println(d)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
