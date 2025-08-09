package main

// Adaptee has a specific API that is incompatible with Target.
type Adaptee struct{}

func (Adaptee) SpecificRequest() string {
	// Returns a reversed string to simulate an incompatible format
	return ".eetpadA eht fo roivaheb laicepS"
}