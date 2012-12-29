package main

import (
	"fmt"
	"github.com/prinsmike/passgo"
)

func main() {

	c := []byte("bcdfghjklmnpqrstvwxyz")
	v := []byte("aeiou")
	n := []byte("0123456789")
	s := []byte("!@$#%&*-_.")
	gen := passgo.NewGenerator(c, v, n, s, true, 8)
	pass, _ := gen.NewPassword(10, 2, 1)
	fmt.Println(pass)
}
