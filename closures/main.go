package main

import (
	"fmt"
)

func main() {

	var afn areaFn = areaFnImpl()

	r := Rectangle{length: 5, width: 10}
	fmt.Printf("%d\n", afn(r.length, r.width))

	s := Square{length: 5, width: 5}
	fmt.Printf("%d\n", afn(s.length, s.width))

	var a adder = adderImpl()
	var m multiplier = multiplierImpl()

	fmt.Println(a(10, 2))
	fmt.Println(m(10, 2))

}

type Rectangle struct {
	length int
	width  int
}

type Square struct {
	length int
	width  int
}

type areaFn func(int, int) int

func areaFnImpl() func(int, int) int {
	return func(length int, width int) int {
		return length * width
	}
}

type adder func(int, int) int

func adderImpl() func(int, int) int {
	return func(n1 int, n2 int) int {
		return n1 + n2
	}

}

type multiplier func(int, int) int

func multiplierImpl() func(int, int) int {
	return func(n1 int, n2 int) int {
		return n1 * n2
	}

}
