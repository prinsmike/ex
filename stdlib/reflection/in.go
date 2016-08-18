package main

import (
	"fmt"
	"reflect"
)

func Index(list interface{}, v interface{}) int {
	return IndexFunc(list, v, reflect.DeepEqual)
}

type EqFunc func(a interface{}, b interface{}) bool

func IndexFunc(list interface{}, v interface{}, eq EqFunc) int {
	l := reflect.ValueOf(list)
	for i := 0; i < l.Len(); i += 1 {
		if eq(v, l.Index(i).Interface()) {
			return i
		}
	}
	return -1
}

func main() {
	l := []int{1, 2, 3, 4, 5, 6}
	for _, v := range []int{5, 7} {
		fmt.Println(v, Index(l, v))
	}
}
