package main

import (
	"fmt"
	"math"
)

type A []float64

func main() {
	a := A{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("a:", a)
	fmt.Println("Product of a:", a.prod())
	fmt.Println("Geometric mean of a:", a.geoMean())
	fmt.Println("Arithmetic mean of a:", a.arithMean())

	fmt.Println("")
	fmt.Println("")
	b := A{1, 2, 3, 4, 5, 6, 7, 8, 9, 1000}
	fmt.Println("b:", b)
	fmt.Println("Product of b:", b.prod())
	fmt.Println("Geometric mean of b:", b.geoMean())
	fmt.Println("Arithmetic mean of b:", b.arithMean())
}

func (a A) prod() float64 {
	p := 1.0
	for _, i := range a {
		p = p * i
	}
	return p
}

func (a A) sum() float64 {
	s := 0.0
	for _, i := range a {
		s = s + i
	}
	return s
}

func (a A) geoMean() float64 {
	return math.Pow(
		a.prod(),
		(1.0 / float64(
			len(a),
		)),
	)
}

func (a A) arithMean() float64 {
	return a.sum() / float64(len(a))
}
