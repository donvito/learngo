package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var channel = make(chan string)

func main() {

	wg.Add(1)
	go hello()
	go world()
	wg.Wait()
}

func hello() {
	str := "HELLO "
	for i := 0; i < 100; i++ {
		str += "HELLO "
	}
	channel <- str
	wg.Done()
}

func world() {
	fmt.Printf("world")
	for {
		fmt.Printf("looping")
		str := <-channel
		fmt.Printf(str)
	}

}
