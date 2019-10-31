package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var s interface{}
	s = "This is a string"
	fmt.Printf("%v\n", reflect.TypeOf(s))

	x := strings.Contains(s.(string), "string")

	println(x)
}
