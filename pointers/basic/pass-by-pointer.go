package main

import (
	"fmt"
)

func main() {
	var x = 1

	fmt.Printf("%d\n", x)
	printx(x)
	fmt.Printf("%d\n", x)
	changex(&x)
	fmt.Printf("%d\n", x)

}

func printx(x int) {
	x = 0
}

func changex(x *int) {
	*x = 2
}
