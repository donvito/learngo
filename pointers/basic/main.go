package main

import (
	"fmt"
	"strings"
)

func main() {
	//declare some variables and add a value to each
	var firstname = "melvin"
	var surname = "vivas"

	//shorthand variable declaration can also be used
	//firstname := "melvin"
	//surname := "vivas"

	//create ptrFirstname to hold the pointer to firstname's memory address
	ptrFirstname := &firstname

	//create ptrSurname to hold the pointer to surname's memory address
	ptrSurname := &surname

	//defrencing - retrieve the actual values using the *pointername notation and convert to uppercase, then concatenate to build the fullname
	fullname := strings.ToUpper(*ptrFirstname) + " " + strings.ToUpper(*ptrSurname)

	fmt.Println(fullname)

}
