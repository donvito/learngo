package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var ch = make(chan int)

func main() {

	runtime.GOMAXPROCS(1)
	fmt.Printf("gomaxprocs = %v", runtime.GOMAXPROCS(-1))
	fmt.Println("\nSTART")

	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("i = %v\n", i)
			} else {
				break
			}
		}

		wg.Done()

	}(ch)

	go func(ch chan<- int) {
		ch <- 8
		ch <- 12
		ch <- 14
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
	fmt.Println("\nEND")

}
