package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var s interface{}
	// s = "This is a string"
	// fmt.Printf("%v", reflect.TypeOf(s))

	// x := strings.Contains(s.(string), "string")

	// println(x)

	m := newMultiplier()
	fmt.Printf("%v", reflect.TypeOf(m))

	print(m(5, 10))

	print(multiply(5, 10))

}

type multiplier func(int, int) int

func newMultiplier() multiplier {
	return func(x int, y int) int {
		product := x * y
		return product
	}

}

func multiply(m multiplier) int {

}
