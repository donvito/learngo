package main

// Target is the interface expected by the client.
type Target interface {
	Request() string
}