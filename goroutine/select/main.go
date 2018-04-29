package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var logCh = make(chan string, 10)
var doneCh = make(chan struct{})

func main() {

	runtime.GOMAXPROCS(1)
	go logger()
	logCh <- "this is string 1"
	logCh <- "this is string 2"
	logCh <- "this is string 3"
	logCh <- "this is string 4"
	logCh <- "this is string 5"
	logCh <- "this is string 6"
	logCh <- "this is string 7"
	logCh <- "this is string 8"
	time.Sleep(5 * time.Second)

	doneCh <- struct{}{}
}

func logger() {
	// for entry := range logCh {
	// 	fmt.Println(entry)
	// }
	for {
		select {
		case entry := <-logCh:
			fmt.Println(entry)
		case <-doneCh:
			break
		}
	}
}
