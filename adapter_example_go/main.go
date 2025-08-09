package main

import "fmt"

func clientCode(t Target) {
	fmt.Println(t.Request())
}

func main() {
	fmt.Println("Client: The Adaptee has a weird interface:")
	adaptee := Adaptee{}
	fmt.Printf("Adaptee: %s\n", adaptee.SpecificRequest())

	fmt.Println("\nClient: But with an adapter, I can work with it:")
	adapter := NewAdapter(adaptee)
	clientCode(adapter)
}