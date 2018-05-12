//https://play.golang.org/p/lRz2A0BF5sm

package main

import (
	"fmt"
)

type Person struct {
	FirstName, LastName string
}

func (p Person) Fullname() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) MutateValue(firstname, lastname string) {
	p.FirstName = firstname
	p.LastName = lastname
}

func (p *Person) MutatePointer(firstname, lastname string) {
	p.FirstName = firstname
	p.LastName = lastname
}

func main() {
	p := Person{"Melvin", "Vivas"}
	p.MutateValue("Aivan", "Monceller")
	//does not mutate
	fmt.Println(p.Fullname())

	p.MutatePointer("Aivan", "Monceller")
	//mutates since receiver is pointer
	fmt.Println(p.Fullname())
}
