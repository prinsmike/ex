package main

import "fmt"

type Common interface {
	String() string
	New() Common
}

type A struct {
	i int
}

func (a *A) String() string { return fmt.Sprintf("A(%v)", a.i) }
func (a *A) New() Common    { return new(A) }

type B struct {
	s string
}

func (b *B) String() string { return fmt.Sprintf("B(%v)", b.s) }
func (b *B) New() Common    { return new(B) }

type C struct {
	f float64
}

func (c *C) String() string { return fmt.Sprintf("C(%v)", c.f) }
func (c *C) New() Common    { return new(C) }

type m map[string]Common

func main() {
	m := map[string]Common{"A": &A{1}, "B": &B{"2"}, "C": &C{3}}
	m1 := make(map[string]Common)
	for i, x := range m {
		m1[i] = x.New()
	}
	fmt.Printf("old %s\n", m)
	fmt.Printf("new %s\n", m1)
}
