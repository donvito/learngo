package main

// Adapter wraps an Adaptee and makes it compatible with Target.
type Adapter struct {
	adaptee Adaptee
}

func NewAdapter(adaptee Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
	translated := reverse(a.adaptee.SpecificRequest())
	return "Adapter: (TRANSLATED) " + translated
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}