package main

import (
	"fmt"
	"reflect"
)

func main() {
	r := Rectangle{length: 5, width: 10}
	printArea(&r)

	s := Square{length: 5, width: 5}
	printArea(&s)
}

type Shape interface {
	area() int
}

type Rectangle struct {
	length int
	width  int
}

type Square struct {
	length int
	width  int
}

func (rect *Rectangle) area() int {
	return rect.length * rect.width
}

func (square *Square) area() int {
	return square.length * square.width
}

func printArea(shape Shape) {
	fmt.Printf("Area of %s is %d \n", reflect.TypeOf(shape), shape.area())
}
