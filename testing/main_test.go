package main

import "testing"

func TestSomething(t *testing.T) {

	s := "this is a string"

	if something(s) != s {
		t.Errorf("error in test")
	}
}
