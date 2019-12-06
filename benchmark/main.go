package main

import (
	"fmt"
)

func sum(x int) int {
	if x <= 0 {
		return 0
	} else {
		return x + sum(x-1)
	}
}

func main() {
	y := sum(10000000)
	fmt.Println(y)
}
