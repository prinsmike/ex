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
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, more := frames.Next()
	fmt.Printf("%s:%d %s - More: %t\n", frame.File, frame.Line, frame.Function, more)
	frame, more = frames.Next()
	fmt.Printf("%s:%d %s - More: %t\n", frame.File, frame.Line, frame.Function, more)
	frame, more = frames.Next()
	fmt.Printf("%s:%d %s - More: %t\n", frame.File, frame.Line, frame.Function, more)
	frame, _ = frames.Next()
	fmt.Printf("%s:%d %s - More: %t\n", frame.File, frame.Line, frame.Function, more)
	frame, _ = frames.Next()
	fmt.Printf("%s:%d %s - More: %t\n", frame.File, frame.Line, frame.Function, more)

}

func myCaller() {
	trace2()
}
