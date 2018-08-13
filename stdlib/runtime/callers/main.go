package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Compiler version: %s\n", runtime.Version())
	trace2()

	myCaller()
}

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func myCaller() {
	trace2()
}
