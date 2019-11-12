package main

import (
	"fmt"
)

func main() {

	//Declare with Type
	var s string = "Melvin"

	fmt.Println(s)

	//Inferred Type using the shorthand notation
	z := "Z"
	fmt.Println(z)

	//Declare an array of strings [no. of elements]Type using var
	var arr1 [5]string
	arr1 = [5]string{"one", "two", "three", "four", "five"}
	fmt.Println(arr1)

	//Declare array using shorthand notation
	arr2 := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println(arr2)

	for i, v := range arr1 {
		fmt.Printf("i =  %v, v = %v\n", i, v)
	}

	//Declare slice, similar to array but no. of elements is not defined
	var slice1 []string
	slice1 = []string{"one", "two"}
	fmt.Println(slice1)

	slice2 := []string{"one", "two", "three", "four", "five"}
	fmt.Println(slice2)

	for i, v := range slice2 {
		fmt.Printf("i =  %v, v = %v\n", i, v)
	}

	//print elements in index 1-3, this is a new slice
	fmt.Println(slice2[1:4])

	//print elements from start to index 3, this is a new slice
	fmt.Println(slice2[:4])

	//print elements from index 0 to end of slice, this is a new slice
	fmt.Println(slice2[0:])

	//declare a map of int and having string as key
	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1
	m["two"] = 2

	delete(m, "one")
	fmt.Println(m)

	//declare a map using shorthand notation
	m1 := make(map[string]int)
	m1 = make(map[string]int)
	m1["one"] = 1
	m1["two"] = 2
	fmt.Println(m1)

	for k, v := range m1 {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}

}
