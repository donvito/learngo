package main

import (
	"fmt"
)

/*
Person Struct
*/
type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "Melvin Vivas", age: 41}
	printPerson(p)
	fmt.Printf("%s, %d\n", p.name, p.age)
	changePerson(&p)
	fmt.Printf("%s, %d\n", p.name, p.age)

}

func printPerson(p Person) {
	fmt.Printf("%s, %d\n", p.name, p.age)
}

func changePerson(p *Person) {
	p.name = "Melvin Dave"
	p.age = 28
}
