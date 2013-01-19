package main

type Test struct {
	S  string
	N  int
	SS []string
}

func main() {

	t := &Test{"Test", 7, []string{"This", "is", "a", "test."}}

	println("Hello, World!")
	println("Hello,", "World!")
	println(t)
}
