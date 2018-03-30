package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	runtime.GOMAXPROCS(1)
	fmt.Printf("gomaxprocs = %v", runtime.GOMAXPROCS(-1))
	fmt.Println("\nSTART")

	for i := 0; i < 100; i++ {
		wg.Add(2)
		m.RLock()
		go sayCount()
		m.Lock()
		go increment()
	}

	wg.Wait()
	fmt.Println("END")

}

func sayCount() {
	fmt.Printf("Counter %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
